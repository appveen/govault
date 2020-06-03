#!/bin/sh
set -e
if [ -f $WORKSPACE/../TOGGLE ]; then
    echo "****************************************************"
    echo "odp:govault :: Toggle mode is on, terminating build"
    echo "odp:govault :: BUILD CANCLED"
    echo "****************************************************"
    exit 0
fi

cd $WORKSPACE

cDate=`date +%Y.%m.%d.%H.%M` #Current date and time

if [ -f $WORKSPACE/../ODP_RELEASE ]; then
    REL=`cat $WORKSPACE/../ODP_RELEASE`
fi
BRANCH='dev'
if [ -f $WORKSPACE/../BRANCH ]; then
    BRANCH=`cat $WORKSPACE/../BRANCH`
fi
if [ $1 ]; then
    REL=$1
fi
if [ ! $REL ]; then
    echo "****************************************************"
    echo "odp:govault :: Please Create file ODP_RELEASE with the releaese at $WORKSPACE or provide it as 1st argument of this script."
    echo "odp:govault :: BUILD FAILED"
    echo "****************************************************"
    exit 0
fi
TAG=$REL

echo "****************************************************"
echo "odp:govault :: Using build :: "$TAG
echo "****************************************************"

echo "****************************************************"
echo "odp:govault :: Adding IMAGE_TAG in Dockerfile :: "$TAG
echo "****************************************************"
sed -i.bak s#__image_tag__#$TAG# Dockerfile

if [ -f $WORKSPACE/../CLEAN_BUILD_GOVAULT ]; then
    echo "****************************************************"
    echo "odp:govault :: Doing a clean build"
    echo "****************************************************"

    docker build --no-cache -t odp:govault.$TAG .
    rm $WORKSPACE/../CLEAN_BUILD_GOVAULT

else
    echo "****************************************************"
    echo "odp:govault :: Doing a normal build"   
    echo "****************************************************"
    docker build -t odp:govault.$TAG .
fi
echo "****************************************************"
echo "odp:govault :: BUILD SUCCESS :: odp:govault.$TAG"
echo "****************************************************"
echo $TAG > $WORKSPACE/../LATEST_GOVAULT
