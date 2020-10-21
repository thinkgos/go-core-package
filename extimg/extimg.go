// Package extimg extend image
package extimg

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

var ext = []string{
	"ase", "art", "bmp", "blp", "cd5", "cit", "cpt", "cr2", "cut", "dds", "dib",
	"djvu", "egt", "exif", "gif", "gpl", "grf", "icns", "ico", "iff", "jng", "jpeg",
	"jpg", "jfif", "jp2", "jps", "lbm", "max", "miff", "mng", "msp", "nitf", "ota",
	"pbm", "pc1", "pc2", "pc3", "pcf", "pcx", "pdn", "pgm", "PI1", "PI2", "PI3",
	"pict", "pct", "pnm", "pns", "ppm", "psb", "psd", "pdd", "psp", "px", "pxm",
	"pxr", "qfx", "raw", "rle", "sct", "sgi", "rgb", "int", "bw", "tga", "tiff",
	"tif", "vtf", "xbm", "xcf", "xpm", "3dv", "amf", "ai", "awg", "cgm", "cdr",
	"cmx", "dxf", "e2d", "egt", "eps", "fs", "gbr", "odg", "svg", "stl", "vrml",
	"x3d", "sxd", "v2d", "vnd", "wmf", "emf", "art", "xar", "png", "webp", "jxr",
	"hdp", "wdp", "cur", "ecw", "iff", "lbm", "liff", "nrrd", "pam", "pcx", "pgf",
	"sgi", "rgb", "rgba", "bw", "int", "inta", "sid", "ras", "sun", "tga",
}

// GetExts get image ext slice
func GetExts() []string {
	return ext
}

// GetType returns the type of image (like image/jpeg)
func GetType(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buf)
	for _, ext := range ext {
		if strings.Contains(ext, filetype[6:]) { // like image/jpeg
			return filetype, nil
		}
	}

	return "", errors.New("invalid image type")
}
