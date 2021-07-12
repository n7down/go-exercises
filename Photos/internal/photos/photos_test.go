package photos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenamePhotos(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"photo.jpg, Warsaw, 2013-09-05 14:08:15\njohn.png, London, 2015-06-20 15:13:22\nmyFriends.png, Warsaw, 2013-09-05 14:07:13\nEiffel.jpg, Paris, 2015-07-23 08:03:02\npisatower.jpg, Paris, 2015-07-22 23:59:59\nBOB.jpg, London, 2015-08-05 00:02:03\nnotredame.png, Paris, 2015-09-01 12:00:00\nme.jpg, Warsaw, 2013-09-06 15:40:22\na.png, Warsaw, 2016-02-13 13:33:50\nb.jpg, Warsaw, 2016-01-02 15:12:22\nc.jpg, Warsaw, 2016-01-02 14:34:30\nd.jpg, Warsaw, 2016-01-02 15:15:01\ne.png, Warsaw, 2016-01-02 09:49:09\nf.png, Warsaw, 2016-01-02 10:55:32\ng.jpg, Warsaw, 2016-02-29 22:13:11", "Warsaw02.jpg\nLondon1.png\nWarsaw01.png\nParis2.jpg\nParis1.jpg\nLondon2.jpg\nParis3.png\nWarsaw03.jpg\nWarsaw09.png\nWarsaw07.jpg\nWarsaw06.jpg\nWarsaw08.jpg\nWarsaw04.png\nWarsaw05.png\nWarsaw10.jpg\n"},
	}

	for _, test := range tests {
		actual := RenamePhotos(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%s should return %s", test.input, test.expected))
	}
}
