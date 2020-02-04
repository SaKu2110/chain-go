package config

import(
	"testing"
)

// 成功例
func TestCheckElementsSuccess(t *testing.T) {
	c = dataBaseConfig {
		User: "admin",
		Pass: "adminpass",
		IP	: "127.0.0.1",
		Port: "3306",
		Name: "test",
	}

    err := checkElements(c)
    if err != nil {
        t.Fatalf("failed test %#v", err)
    }
}

// Userの値が存在しない
func TestCheckElementsFaildUserValue(t *testing.T) {
	c = dataBaseConfig {
		User: "",
		Pass: "adminpass",
		IP	: "127.0.0.1",
		Port: "3306",
		Name: "test",
	}

    err := checkElements(c)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

func TestCheckElementsFaildPassValue(t *testing.T) {
	c = dataBaseConfig {
		User: "admin",
		Pass: "",
		IP	: "127.0.0.1",
		Port: "3306",
		Name: "test",
	}

    err := checkElements(c)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

func TestCheckElementsFaildIPValue(t *testing.T) {
	c = dataBaseConfig {
		User: "admin",
		Pass: "adminpass",
		IP	: "",
		Port: "3306",
		Name: "test",
	}

    err := checkElements(c)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

func TestCheckElementsFaildPortValue(t *testing.T) {
	c = dataBaseConfig {
		User: "admin",
		Pass: "adminpass",
		IP	: "127.0.0.1",
		Port: "",
		Name: "test",
	}

    err := checkElements(c)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}

func TestCheckElementsFaildNameValue(t *testing.T) {
	c = dataBaseConfig {
		User: "admin",
		Pass: "adminpass",
		IP	: "127.0.0.1",
		Port: "3306",
		Name: "",
	}

    err := checkElements(c)
    if err == nil {
        t.Fatalf("failed test %#v", err)
    }
}