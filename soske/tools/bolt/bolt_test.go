package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/Soske/soske/tools/test"
)

func TestDeleteBucket(t *testing.T) {
	// setup
	db := test.MakeDB()
	defer test.KillDB(db)

	// success
	err := DeleteBucket(db, "alpha")
	test.AssertNoBucket(t, db, "alpha")
	assert.NoError(t, err)
}

func TestGetBucket(t *testing.T) {
	// setup
	db := test.MakeDB()
	defer test.KillDB(db)

	// success
	bmap, err := GetBucket(db, "alpha")
	assert.Equal(t, test.TestData["alpha"], bmap)
	assert.NoError(t, err)
}

func TestGetKey(t *testing.T) {
	// setup
	db := test.MakeDB()
	defer test.KillDB(db)

	// success
	val, err := GetKey(db, "alpha", "body")
	assert.Equal(t, []byte("Alpha value."), val)
	assert.NoError(t, err)
}

func TestListBuckets(t *testing.T) {
	// setup
	db := test.MakeDB()
	defer test.KillDB(db)

	// success
	names, err := ListBuckets(db)
	assert.Equal(t, []string{"alpha", "bravo"}, names)
	assert.NoError(t, err)
}

func TestSetBucket(t *testing.T) {
	// setup
	db := test.MakeDB()
	defer test.KillDB(db)

	// success
	err := SetBucket(db, "test", test.TestData["alpha"])
	test.AssertBucket(t, db, "test", test.TestData["alpha"])
	assert.NoError(t, err)
}
