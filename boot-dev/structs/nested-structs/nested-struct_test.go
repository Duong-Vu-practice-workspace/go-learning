package nested_structs

import "testing"

func TestCanSendMessage_AllFieldsPresent(t *testing.T) {
	m := messageToSend{
		message: "Hi",
		sender:  user{name: "Alice", number: 1111},
		recipient: user{
			name:   "Bob",
			number: 2222,
		},
	}
	if !canSendMessage(m) {
		t.Fatalf("canSendMessage returned false; want true")
	}
}

func TestCanSendMessage_MissingSenderName(t *testing.T) {
	m := messageToSend{
		message: "Hi",
		sender:  user{name: "", number: 1111},
		recipient: user{
			name:   "Bob",
			number: 2222,
		},
	}
	if canSendMessage(m) {
		t.Fatalf("canSendMessage returned true for missing sender name; want false")
	}
}

func TestCanSendMessage_MissingSenderNumber(t *testing.T) {
	m := messageToSend{
		message: "Hi",
		sender:  user{name: "Alice", number: 0},
		recipient: user{
			name:   "Bob",
			number: 2222,
		},
	}
	if canSendMessage(m) {
		t.Fatalf("canSendMessage returned true for missing sender number; want false")
	}
}

func TestCanSendMessage_MissingRecipientName(t *testing.T) {
	m := messageToSend{
		message: "Hi",
		sender:  user{name: "Alice", number: 1111},
		recipient: user{
			name:   "",
			number: 2222,
		},
	}
	if canSendMessage(m) {
		t.Fatalf("canSendMessage returned true for missing recipient name; want false")
	}
}

func TestCanSendMessage_MissingRecipientNumber(t *testing.T) {
	m := messageToSend{
		message: "Hi",
		sender:  user{name: "Alice", number: 1111},
		recipient: user{
			name:   "Bob",
			number: 0,
		},
	}
	if canSendMessage(m) {
		t.Fatalf("canSendMessage returned true for missing recipient number; want false")
	}
}

func TestCanSendMessage_ZeroValueStruct(t *testing.T) {
	var m messageToSend
	if canSendMessage(m) {
		t.Fatalf("canSendMessage returned true for zero-value struct; want false")
	}
}
