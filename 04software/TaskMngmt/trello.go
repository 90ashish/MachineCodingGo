package trello

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type Card struct {
	ID          string
	Name        string
	Description string
	AssignedTo  *User
}

type List struct {
	ID    string
	Name  string
	Cards map[string]*Card
}

type Board struct {
	ID      string
	Name    string
	Privacy string
	URL     string
	Members map[string]*User
	Lists   map[string]*List
}

type App struct {
	Users  map[string]*User
	Boards map[string]*Board
}

func NewApp() *App {
	return &App{
		Users:  make(map[string]*User),
		Boards: make(map[string]*Board),
	}
}

func generateID() string {
	return uuid.New().String()
}

func (a *App) CreateUser(name, email string) *User {
	user := &User{
		ID:    generateID(),
		Name:  name,
		Email: email,
	}
	a.Users[user.ID] = user
	fmt.Printf("Created User: %s\n%s\n", user.ID, user.Name)
	return user
}

func (a *App) CreateBoard(name, privacy string) *Board {
	if privacy == "" {
		privacy = "PUBLIC"
	}

	projectID := generateID()

	board := &Board{
		ID:      projectID,
		Name:    name,
		Privacy: privacy,
		URL:     "http://localhost/" + projectID,
		Members: make(map[string]*User),
		Lists:   make(map[string]*List),
	}
	a.Boards[board.ID] = board
	fmt.Printf("Created Board: %s\n%s\n", board.ID, board.Name)
	return board
}

func (b *Board) CreateList(name string) *List {
	list := &List{
		ID:    generateID(),
		Name:  name,
		Cards: make(map[string]*Card),
	}
	b.Lists[list.ID] = list
	fmt.Printf("Created List: %s\n%s\n", list.ID, list.Name)
	return list
}

func (l *List) CreateCard(name, description string) *Card {
	card := &Card{
		ID:          generateID(),
		Name:        name,
		Description: description,
	}
	l.Cards[card.ID] = card
	fmt.Printf("Created Card: %s\n%s\n", card.ID, card.Name)
	return card
}

func (a *App) DeleteBoard(boardID string) {
	delete(a.Boards, boardID)
}

func (b *Board) AddMember(user *User) {
	b.Members[user.ID] = user
}

func (b *Board) RemoveMember(userID string) {
	delete(b.Members, userID)
}

func (b *Board) DeleteList(listID string) {
	delete(b.Lists, listID)
}

func (l *List) DeleteCard(cardID string) {
	delete(l.Cards, cardID)
}

func (c *Card) AssignTo(user *User) {
	c.AssignedTo = user
}

func (c *Card) Unassign() {
	c.AssignedTo = nil
}

func (l *List) MoveCard(cardID string, destList *List) {
	card, exists := l.Cards[cardID]
	if !exists {
		fmt.Printf("Card %s does not exist\n", cardID)
		return
	}
	delete(l.Cards, cardID)
	destList.Cards[cardID] = card
}

func (a *App) ShowAllBoards() {
	if len(a.Boards) == 0 {
		fmt.Println("No boards")
		return
	}
	for _, board := range a.Boards {
		fmt.Printf("%+v\n", *board)
	}
}

func (a *App) ShowBoard(boardID string) {
	board, exists := a.Boards[boardID]
	if !exists {
		fmt.Printf("Board %s does not exist\n", boardID)
		return
	}
	fmt.Printf("%+v\n", *board)
}

func (b *Board) ShowList(listID string) {
	list, exists := b.Lists[listID]
	if !exists {
		fmt.Printf("List %s does not exist\n", listID)
		return
	}
	fmt.Printf("%+v\n", *list)
}

func (b *Board) ShowCard(listID, cardID string) {
	list, exists := b.Lists[listID]
	if !exists {
		fmt.Printf("List %s does not exist\n", listID)
		return
	}
	card, exists := list.Cards[cardID]
	if !exists {
		fmt.Printf("Card %s does not exist\n", cardID)
		return
	}
	fmt.Printf("%+v\n", *card)
}

func Trello() {
	system := NewApp()

	// Create users
	user1 := system.CreateUser("Alice", "alice@example.com")
	user2 := system.CreateUser("Bob", "bob@example.com")

	// Create a board
	board := system.CreateBoard("Work", "PUBLIC")

	// Add members to the board
	board.AddMember(user1)
	board.AddMember(user2)

	// Create lists
	list1 := board.CreateList("To Do")
	list2 := board.CreateList("In Progress")

	// Create cards
	card1 := list1.CreateCard("Task 1", "Description of Task 1")
	card2 := list1.CreateCard("Task 2", "Description of Task 2")

	// Assign user to card
	card1.AssignTo(user1)
	card2.AssignTo(user2)

	// Move card to another list
	list1.MoveCard(card1.ID, list2)

	// Unassign user from card
	card1.Unassign()

	// Show all boards
	system.ShowAllBoards()

	// Show specific board
	system.ShowBoard(board.ID)

	// Show specific list
	board.ShowList(list1.ID)

	// Show specific card
	board.ShowCard(list2.ID, card1.ID)
}
