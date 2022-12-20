// Copyright © 2020. All rights reserved.
// Author: Ilya Stroy.
// Contacts: iyuryevich@pm.me, https://github.com/qioalice
// License: https://opensource.org/licenses/MIT

package ekaclike_test

import (
	"testing"

	"github.com/qioalice/ekago_private/ekaclike/v4"

	"github.com/stretchr/testify/assert"
)

func TestUnpackInterface(t *testing.T) {

	type T1 struct{}
	type T2 struct{}

	var t1 = T1{}
	var t2 = T2{}

	assert.Equal(t, ekaclike.UnpackInterface(&t1).Type, ekaclike.UnpackInterface(new(T1)).Type)
	assert.Equal(t, ekaclike.UnpackInterface(&t2).Type, ekaclike.UnpackInterface(new(T2)).Type)

	assert.NotEqual(t, ekaclike.UnpackInterface(&t1).Type, ekaclike.UnpackInterface(new(T2)).Type)
	assert.NotEqual(t, ekaclike.UnpackInterface(&t2).Type, ekaclike.UnpackInterface(new(T1)).Type)

	assert.NotEqual(t, ekaclike.UnpackInterface(t1).Type, ekaclike.UnpackInterface(t2).Type)
}
