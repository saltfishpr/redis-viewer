#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

go mod tidy
mkdir -p "$WORKDIR"/output/bin/ && go build -o "$WORKDIR"/output/bin/"$APP_NAME"

echo "#!/bin/bash" >"$WORKDIR"/output/run.sh
echo "LANG=en_US:UTF-8 LANGUAGE=en_US:en ${WORKDIR}/output/bin/${APP_NAME}" >>"$WORKDIR"/output/run.sh
