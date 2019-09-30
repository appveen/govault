package vault

import (
	"errors"
	"sync"
)

//Vault - Base Vault structure
type Vault struct {
	mutex         sync.Mutex
	storePassword string
	trustore      *DB
}

//CreateVault - create new vault with specified location and password
func CreateVault(filePath string, newPassword string) *Vault {
	bucketName := DEFAULTBUCKETNAME
	DB := InitDB(filePath, bucketName)
	random := RandomString(RANDOMSTRINGLENGTH)
	e1, _ := Encrypt([]byte(random), newPassword)
	DB.Upsert(DEFAULTRANDOMKEYNAME, e1)
	e2, _ := Encrypt([]byte(newPassword), random)
	DB.Upsert(DEFAULTPASSWORDKEYNAME, e2)
	return &Vault{
		trustore:      DB,
		storePassword: newPassword,
	}
}

//InitVault - Initialize access to vault
func InitVault(filePath string, storePassword string) (*Vault, error) {
	bucketName := DEFAULTBUCKETNAME
	DB := InitDB(filePath, bucketName)
	encyptedRandomKey, err := DB.Get(DEFAULTRANDOMKEYNAME)
	if err != nil {
		return nil, err
	}
	decryptedRandomKey, err := Decrypt(encyptedRandomKey, storePassword)
	if err != nil {
		return nil, err
	}
	encryptedPassword, err := DB.Get(DEFAULTPASSWORDKEYNAME)
	if err != nil {
		return nil, err
	}
	extractedPassword, err := Decrypt(encryptedPassword, string(decryptedRandomKey))
	if err != nil {
		return nil, err
	}
	if string(extractedPassword) != storePassword {
		return nil, errors.New("Incorrect Password")
	}
	return &Vault{
		trustore:      DB,
		storePassword: storePassword,
	}, nil
}

//Get - get data from trustore
func (v *Vault) Get(key string) ([]byte, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	encyptedRandomKey, err := v.trustore.Get(DEFAULTRANDOMKEYNAME)
	if err != nil {
		return nil, err
	}
	cipheredData, err := v.trustore.Get(key)
	if err != nil {
		return nil, err
	}
	if cipheredData == nil {
		return nil, nil
	}
	return Decrypt(cipheredData, string(encyptedRandomKey))
}

//Upsert - add/update value in trustore
func (v *Vault) Upsert(key string, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	encyptedRandomKey, err := v.trustore.Get(DEFAULTRANDOMKEYNAME)
	if err != nil {
		return err
	}
	e, err := Encrypt([]byte(value), string(encyptedRandomKey))
	if err != nil {
		return err
	}
	err = v.trustore.Upsert(key, e)
	if err != nil {
		return err
	}
	return nil
}

//Delete - delete a value by key in trustore
func (v *Vault) Delete(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.trustore.Delete(key)
	return err
}

//ChangePassword - change trustore password
func (v *Vault) ChangePassword(newPassword string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	encyptedRandomKey, err := v.trustore.Get(DEFAULTRANDOMKEYNAME)
	if err != nil {
		return err
	}
	tx, err := v.trustore.STORE.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	bucket := tx.Bucket([]byte(DEFAULTBUCKETNAME))
	cursor := bucket.Cursor()
	data := map[string]string{}
	for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
		switch string(k) {
		case DEFAULTRANDOMKEYNAME:
			break
		case DEFAULTPASSWORDKEYNAME:
			break
		default:
			d, err := Decrypt(v, string(encyptedRandomKey))
			if err != nil {
				return err
			}
			data[string(k)] = string(d)
			break
		}
	}
	decryptedRandomKey, err := Decrypt(encyptedRandomKey, v.storePassword)
	if err != nil {
		return err
	}
	e, err := Encrypt(decryptedRandomKey, newPassword)
	if err != nil {
		return err
	}
	newEncryptedRandomKey := string(e)
	for key, value := range data {
		switch key {
		case DEFAULTRANDOMKEYNAME:
			break
		case DEFAULTPASSWORDKEYNAME:
			break
		default:
			e, err := Encrypt([]byte(value), newEncryptedRandomKey)
			if err != nil {
				return err
			}
			err = bucket.Put([]byte(key), e)
			if err != nil {
				return err
			}
			break
		}
	}
	err = bucket.Put([]byte(DEFAULTRANDOMKEYNAME), []byte(newEncryptedRandomKey))
	if err != nil {
		return err
	}
	e2, err := Encrypt([]byte(newPassword), string(decryptedRandomKey))
	if err != nil {
		return err
	}
	err = bucket.Put([]byte(DEFAULTPASSWORDKEYNAME), e2)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

//Close - Close trustore
func (v *Vault) Close() error {
	err := v.trustore.CloseDB()
	return err
}
