module github.com/conejoninja/badger2040

go 1.22

require (
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	tinygo.org/x/drivers v0.28.0
	tinygo.org/x/tinydraw v0.4.0
	tinygo.org/x/tinyfont v0.4.0
)

require github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect

replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240515082542-5f2645f5444d
