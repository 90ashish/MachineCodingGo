
### 1. **Code Visual Flow (Diagram)**

```
+------------------+
|      App         |
+------------------+
| Users            |<-----------------------+
| Boards           |                        |
+------------------+                        |
        |                                   |
        v                                   |
+------------------+                        |
|    Board         |                        |
+------------------+                        |
| ID               |                        |
| Name             |                        |
| Privacy          |                        |
| URL              |                        |
| Members          |<---------------------+ |
| Lists            |                      | |
+------------------+                      | |
        |                                 | |
        v                                 | |
+------------------+                      | |
|     List         |                      | |
+------------------+                      | |
| ID               |                      | |
| Name             |                      | |
| Cards            |<-------------------+ | |
+------------------+                    | | |
        |                               | | |
        v                               | | |
+------------------+                    | | |
|     Card         |                    | | |
+------------------+                    | | |
| ID               |                    | | |
| Name             |                    | | |
| Description      |                    | | |
| AssignedTo       |<----+              | | |
+------------------+     |              | | |
                         |              | | |
+------------------+     |              | | |
|    User          |<----+              | | |
+------------------+                    | | |
| ID               |                    | | |
| Name             |                    | | |
| Email            |                    | | |
+------------------+                    | | |
                                        | | |
+-------------------------------------+ | | |
|                Operations           | | | |
+-------------------------------------+ | | |
| 1. CreateUser(name, email)          | | | |
| 2. CreateBoard(name, privacy)       | | | |
| 3. CreateList(name)                 | | | |
| 4. CreateCard(name, description)    | | | |
| 5. AssignTo(user)                   | | | |
| 6. MoveCard(cardID, destList)       | | | |
| 7. ShowAllBoards()                  | | | |
| 8. ShowBoard(boardID)               | | | |
| 9. ShowList(listID)                 | | | |
| 10. ShowCard(listID, cardID)        | | | |
| 11. AddMember(user)                 | | | |
| 12. RemoveMember(userID)            | | | |
| 13. DeleteBoard(boardID)            | | | |
| 14. DeleteList(listID)              | | | |
| 15. DeleteCard(cardID)              | | | |
| 16. Unassign()                      | | | |
+-------------------------------------+ | | |
                                        | | |
                                        | | |
                                        | | |
+-------------------------------------+ | | |
|          Data Flow Summary          | | | |
+-------------------------------------+ | | |
| 1. App holds Users and Boards       | | | |
| 2. Board holds Lists and Members    | | | |
| 3. List holds Cards                 | | | |
| 4. Card can be assigned to a User   | | | |
+-------------------------------------+ | | |
                                        | | |
+-----------------------------------------+ | 
|  Operations take place based on:          |
|  - User interactions                      |
|  - Board interactions                     |
|  - List interactions                      |
|  - Card interactions                      |
+-------------------------------------------+
```

### 2. **Visualization Table (Relationships and Entities)**

This table outlines the key entities and their relationships in the system:

| **Entity** | **Attributes** | **Relationship** | **Description** |
|------------|----------------|------------------|-----------------|
| `App`      | `Users`, `Boards` | Contains `User` and `Board` objects | Main container for all data |
| `User`     | `ID`, `Name`, `Email` | Members of `Board` | Represents a user in the system |
| `Board`    | `ID`, `Name`, `Privacy`, `URL`, `Members`, `Lists` | Contains `List` objects and associated `User` objects | Represents a project or workspace |
| `List`     | `ID`, `Name`, `Cards` | Contains `Card` objects | Represents a collection of tasks or cards |
| `Card`     | `ID`, `Name`, `Description`, `AssignedTo` | Can be assigned to a `User` | Represents a task or item in a list |
| **Operation** | **Description** | **Entity** | **Effect** |
| `CreateUser` | Creates a new user | `App` | Adds a new `User` to the `App` |
| `CreateBoard` | Creates a new board | `App` | Adds a new `Board` to the `App` |
| `CreateList` | Creates a new list | `Board` | Adds a new `List` to a `Board` |
| `CreateCard` | Creates a new card | `List` | Adds a new `Card` to a `List` |
| `AssignTo` | Assigns a card to a user | `Card` | Links a `Card` with a `User` |
| `MoveCard` | Moves a card to another list | `List` | Transfers a `Card` between `List` objects |
| `AddMember` | Adds a user to a board | `Board` | Links a `User` to a `Board` |
| `RemoveMember` | Removes a user from a board | `Board` | Removes the link between `User` and `Board` |
| `DeleteBoard` | Deletes a board | `App` | Removes a `Board` from the `App` |
| `DeleteList` | Deletes a list | `Board` | Removes a `List` from a `Board` |
| `DeleteCard` | Deletes a card | `List` | Removes a `Card` from a `List` |
| `Unassign` | Unassigns a card from a user | `Card` | Removes the link between `Card` and `User` |
| `ShowAllBoards` | Displays all boards | `App` | Outputs all `Board` objects |
| `ShowBoard` | Displays a specific board | `App` | Outputs a single `Board` |
| `ShowList` | Displays a specific list | `Board` | Outputs a single `List` |
| `ShowCard` | Displays a specific card | `List` | Outputs a single `Card` |
