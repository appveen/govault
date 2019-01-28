package vault

import (
	"errors"
)

//Vault - Base Vault structure
type Vault struct {
	storePassword string
	trustore      *DB
}

//CreateVault - create new vault with specified location and password
func CreateVault(filePath string, newPassword string) *Vault {
	bucketName := DEFAULTBUCKETNAME
	DB := InitDB(filePath, bucketName)
	random := RandomString(RANDOMSTRINGLENGTH)
	DB.Upsert(DEFAULTRANDOMKEYNAME, Encrypt([]byte(random), newPassword))
	DB.Upsert(DEFAULTPASSWORDKEYNAME, Encrypt([]byte(newPassword), random))
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
	decryptedRandomKey := Decrypt(encyptedRandomKey, storePassword)
	encryptedPassword, err := DB.Get(DEFAULTPASSWORDKEYNAME)
	if err != nil {
		return nil, err
	}
	extractedPassword := Decrypt(encryptedPassword, string(decryptedRandomKey))
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
	return Decrypt(cipheredData, string(encyptedRandomKey)), nil
}

//Upsert - add/update value in trustore
func (v *Vault) Upsert(key string, value string) error {
	encyptedRandomKey, err := v.trustore.Get(DEFAULTRANDOMKEYNAME)
	if err != nil {
		return err
	}
	err = v.trustore.Upsert(key, Encrypt([]byte(value), string(encyptedRandomKey)))
	if err != nil {
		return err
	}
	return nil
}

//Delete - delete a value by key in trustore
func (v *Vault) Delete(key string) error {
	err := v.trustore.Delete(key)
	return err
}

//ChangePassword - change trustore password
func (v *Vault) ChangePassword(newPassword string) error {
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
			data[string(k)] = string(Decrypt(v, string(encyptedRandomKey)))
			break
		}
	}
	decryptedRandomKey := Decrypt(encyptedRandomKey, v.storePassword)
	newEncryptedRandomKey := string(Encrypt(decryptedRandomKey, newPassword))
	for key, value := range data {
		switch key {
		case DEFAULTRANDOMKEYNAME:
			break
		case DEFAULTPASSWORDKEYNAME:
			break
		default:
			err = bucket.Put([]byte(key), Encrypt([]byte(value), newEncryptedRandomKey))
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
	err = bucket.Put([]byte(DEFAULTPASSWORDKEYNAME), Encrypt([]byte(newPassword), string(decryptedRandomKey)))
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
