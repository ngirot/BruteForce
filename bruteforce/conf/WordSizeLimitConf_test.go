package conf

import "testing"

func TestWordSizeLimitConf_EmptyParamShouldHaveNoSizeLimit(t *testing.T) {
	var conf, _ = NewWordSizeLimitConf("")

	if conf.Min != 0 {
		t.Errorf("Expected '%d' but found '%d'", 0, conf.Min)
	}

	if conf.Max != 0 {
		t.Errorf("Expected '%d' but found '%d'", 0, conf.Max)
	}
}

func TestWordSizeLimitConf_SingleNumberShouldBeMinAndMax(t *testing.T) {
	var conf, _ = NewWordSizeLimitConf("10")

	if conf.Min != 10 {
		t.Errorf("Expected '%d' but found '%d'", 10, conf.Min)
	}

	if conf.Max != 10 {
		t.Errorf("Expected '%d' but found '%d'", 10, conf.Max)
	}
}

func TestWordSizeLimitConf_InvalidParamShouldReturnAnError(t *testing.T) {
	var _, err = NewWordSizeLimitConf("test")

	if err == nil {
		t.Errorf("An error should be return with '%s' parameter", "test")
	}
}

func TestWordSizeLimitConf_RangeShouldSetMinAndMax(t *testing.T) {
	var conf, _ = NewWordSizeLimitConf("5-8")

	if conf.Min != 5 {
		t.Errorf("Expected '%d' but found '%d'", 5, conf.Min)
	}

	if conf.Max != 8 {
		t.Errorf("Expected '%d' but found '%d'", 8, conf.Max)
	}
}

func TestWordSizeLimitConf_ShouldReturnErrorOnThreeNumberRange(t *testing.T) {
	var _, err = NewWordSizeLimitConf("5-8-9")
	if err == nil {
		t.Errorf("An error should be return with '%s' parameter", "5-8-9")
	}
}

func TestWordSizeLimitConf_ShouldReturnErrorrangeWithInvalidFirstNumber(t *testing.T) {
	var _, err = NewWordSizeLimitConf("test-10")
	if err == nil {
		t.Errorf("An error should be return with '%s' parameter", "test-10")
	}
}

func TestWordSizeLimitConf_ShouldReturnErrorrangeWithInvalidSecondNumber(t *testing.T) {
	var _, err = NewWordSizeLimitConf("5-val")
	if err == nil {
		t.Errorf("An error should be return with '%s' parameter", "5-val")
	}
}

func TestWordSizeLimitConf_MinAndMaxShouldBeInvertedIfNecessary(t *testing.T) {
	var conf, _ = NewWordSizeLimitConf("8-3")

	if conf.Min != 3 {
		t.Errorf("Expected '%d' but found '%d'", 3, conf.Min)
	}

	if conf.Max != 8 {
		t.Errorf("Expected '%d' but found '%d'", 8, conf.Max)
	}
}
