package basegorm

import (
	"testing"
	"time"
)

type testUserModel struct {
	BaseModel[int]

	Username string
	Age      int
}

func TestCreateModels(t *testing.T) {
	testValues := []struct {
		ID         int
		Username   string
		Age        int
		DateCreate time.Time
		DateUpdate time.Time
		IsDelete   bool
	}{
		{1, "test1", 25, time.Unix(1689428527, 0), time.Unix(1689428527, 0), false},
		{2, "test2", 100, time.Unix(1657892527, 0), time.Unix(1657892527, 0), false},
		{3, "test3", 0, time.Unix(1, 0), time.Unix(1, 1), true},
	}

	testModels := make([]testUserModel, len(testValues))
	for i, values := range testValues {
		model := testUserModel{
			Username: values.Username,
			Age:      values.Age,
		}
		model.SetId(values.ID)
		model.SetDateCreate(values.DateCreate)
		model.SetDateUpdate(values.DateUpdate)
		model.SetIsDelete(values.IsDelete)
		testModels[i] = model
	}

	// check valid
	for i, i2 := range testModels {
		if i2.GetId() != testValues[i].ID {
			t.Errorf("invalid id")
		}
		if i2.Username != testValues[i].Username {
			t.Errorf("invalid username")
		}
		if i2.Age != testValues[i].Age {
			t.Errorf("invalid age")
		}
		if i2.GetDateCreate() != testValues[i].DateCreate {
			t.Errorf("invalid date create")
		}
		if i2.GetDateUpdate() != testValues[i].DateUpdate {
			t.Errorf("invalid date update")
		}
		if i2.GetIsDelete() != testValues[i].IsDelete {
			t.Errorf("invalid is delete")
		}
	}
}
func TestCreatePointerModels(t *testing.T) {
	testValues := []struct {
		ID         int
		Username   string
		Age        int
		DateCreate time.Time
		DateUpdate time.Time
		IsDelete   bool
	}{
		{1, "test1", 25, time.Unix(1689428527, 0), time.Unix(1689428527, 0), false},
		{2, "test2", 100, time.Unix(1657892527, 0), time.Unix(1657892527, 0), false},
		{3, "test3", 0, time.Unix(1, 0), time.Unix(1, 1), true},
	}

	testModels := make([]*testUserModel, len(testValues))
	for i, values := range testValues {
		model := &testUserModel{
			Username: values.Username,
			Age:      values.Age,
		}
		model.SetId(values.ID)
		model.SetDateCreate(values.DateCreate)
		model.SetDateUpdate(values.DateUpdate)
		model.SetIsDelete(values.IsDelete)
		testModels[i] = model
	}

	// check valid
	for i, i2 := range testModels {
		if i2.GetId() != testValues[i].ID {
			t.Errorf("invalid id")
		}
		if i2.Username != testValues[i].Username {
			t.Errorf("invalid username")
		}
		if i2.Age != testValues[i].Age {
			t.Errorf("invalid age")
		}
		if i2.GetDateCreate() != testValues[i].DateCreate {
			t.Errorf("invalid date create")
		}
		if i2.GetDateUpdate() != testValues[i].DateUpdate {
			t.Errorf("invalid date update")
		}
		if i2.GetIsDelete() != testValues[i].IsDelete {
			t.Errorf("invalid is delete")
		}
	}
}

func TestBaseGetTx(t *testing.T) {
	var model testUserModel

	if op := model.TxOptionNotDelete(); op == nil {
		t.Error("op NotDelete return is nil")
	}

	if op := model.TxOptionById(1); op == nil {
		t.Error("op ById return is nil")
	}

	if op := model.TxOptionPreloadNotDelete("test"); op == nil {
		t.Error("op PreloadNotDelete return is nil")
	}
}
