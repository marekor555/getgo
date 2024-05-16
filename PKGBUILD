# Maintainer: MAREKOR555 <marekor555@interia.pl>

pkgname=getgo
pkgver=1.0.0
epoch=1
pkgrel=1
pkgdesc="alternative to curl/wget written in golang"
url="http://github.com/marekor555"
arch=('any')
makedepends=('go')
license=('GPLv3')
provides=('getgo')

package() {
	install -Dm755 getgo ${pkgdir}/usr/bin/getgo
}

build() {
	pwd
	cd ..
	go get .
	go build .
	cp getgo src
}
