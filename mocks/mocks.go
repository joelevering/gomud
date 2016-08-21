// Pausing work here to figure out the issue detailed in room_test.go
// package mocks
//
// type MockRoom struct {
// 	Messages []string
// 	Clients  []*Client
// }
//
// func (m *MockRoom) GetId() int {
// 	return 0
// }
//
// func (m *MockRoom) GetName() string {
// 	return "Name"
// }
//
// func (m *MockRoom) GetDesc() string {
// 	return "Desc"
// }
//
// func (m *MockRoom) GetExits() []Exit {
// 	return []Exit{}
// }
//
// func (m *MockRoom) GetClients() []*Client {
// 	return []*Client{}
// }
//
// func (m *MockRoom) GetNpcs() []NPC {
// 	return []NPC{}
// }
//
// func (m *MockRoom) SetClients(clients []*Client) {
// 	m.Clients = clients
// }
//
// func (m *MockRoom) Message(msg string) {
// 	m.Messages = append(m.Messages, msg)
// }
