#！/bin/bash
## SPDX-License-Identifier: Apache-2.0
## Copyright Authors of Spider

DOWNLOAD_DIR="$1"
[ -z "$DOWNLOAD_DIR" ] && echo "error, miss DOWNLOAD_DIR" && exit 1
mkdir -p $DOWNLOAD_DIR
echo "download package to $DOWNLOAD_DIR "

[ -z "$CNI_PACKAGE_VERSION" ] && echo "error, miss CNI_PACKAGE_VERSION " && exit 1
echo "use CNI_PACKAGE_VERSION $CNI_PACKAGE_VERSION"

[ -z "$IMAGE_MUTLUS" ] && echo "error, miss IMAGE_MUTLUS " && exit 1
echo "use IMAGE_MUTLUS $IMAGE_MUTLUS"

[ -z "$IMAGE_WHEREABOUTS" ] && echo "error, miss IMAGE_WHEREABOUTS " && exit 1
echo "use IMAGE_WHEREABOUTS $IMAGE_WHEREABOUTS"

[ -z "$TEST_IMAGE" ] && echo "error, miss TEST_IMAGE " && exit 1
echo "use TEST_IMAGE $TEST_IMAGE"


#=================================

OS=$(uname | tr 'A-Z' 'a-z')

# prepare cni-plugins
PACKAGE_NAME="cni-plugins-linux-amd64-${CNI_PACKAGE_VERSION}.tgz"
if [ ! -f  "${DOWNLOAD_DIR}/${PACKAGE_NAME}" ]; then
  echo "begin to download cni-plugins ${PACKAGE_NAME} "
  wget -P ${DOWNLOAD_DIR} https://github.com/containernetworking/plugins/releases/download/${CNI_PACKAGE_VERSION}/${PACKAGE_NAME}
else
  echo "${DOWNLOAD_DIR}/${PACKAGE_NAME} exist, skip download..."
fi

#=================================

# prepare image
IMAGES="
$IMAGE_MUTLUS \
$IMAGE_WHEREABOUTS \
$TEST_IMAGE"

for image in $IMAGES
do
  SUFF_IMAGE=$(echo $image | awk -F ':' '{print $1}')
  EXIST=$(docker images | grep $SUFF_IMAGE)
  if test -z "$EXIST"; then
    echo "Image: $image not exist locally, will pull..."
    docker pull $image
  else
    echo "Image: $image already exist locally"
  fi
done