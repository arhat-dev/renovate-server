# Copyright 2020 The arhat.dev Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# build
image.build.renovate-server.linux.x86:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.amd64:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.armv5:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.armv6:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.armv7:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.arm64:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.ppc64le:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.mips64le:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.s390x:
	sh scripts/image/build.sh $@

image.build.renovate-server.linux.all: \
	image.build.renovate-server.linux.amd64 \
	image.build.renovate-server.linux.arm64 \
	image.build.renovate-server.linux.armv7 \
	image.build.renovate-server.linux.armv6 \
	image.build.renovate-server.linux.armv5 \
	image.build.renovate-server.linux.x86 \
	image.build.renovate-server.linux.s390x \
	image.build.renovate-server.linux.ppc64le \
	image.build.renovate-server.linux.mips64le

image.build.renovate-server.windows.amd64:
	sh scripts/image/build.sh $@

image.build.renovate-server.windows.armv7:
	sh scripts/image/build.sh $@

image.build.renovate-server.windows.all: \
	image.build.renovate-server.windows.amd64 \
	image.build.renovate-server.windows.armv7

# push
image.push.renovate-server.linux.x86:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.amd64:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.armv5:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.armv6:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.armv7:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.arm64:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.ppc64le:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.mips64le:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.s390x:
	sh scripts/image/push.sh $@

image.push.renovate-server.linux.all: \
	image.push.renovate-server.linux.amd64 \
	image.push.renovate-server.linux.arm64 \
	image.push.renovate-server.linux.armv7 \
	image.push.renovate-server.linux.armv6 \
	image.push.renovate-server.linux.armv5 \
	image.push.renovate-server.linux.x86 \
	image.push.renovate-server.linux.s390x \
	image.push.renovate-server.linux.ppc64le \
	image.push.renovate-server.linux.mips64le

image.push.renovate-server.windows.amd64:
	sh scripts/image/push.sh $@

image.push.renovate-server.windows.armv7:
	sh scripts/image/push.sh $@

image.push.renovate-server.windows.all: \
	image.push.renovate-server.windows.amd64 \
	image.push.renovate-server.windows.armv7
