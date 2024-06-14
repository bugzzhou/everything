import tkinter as tk
import random

class Game2048:
    def __init__(self, root):
        self.root = root
        self.root.title("2048游戏")
        self.board = [[0] * 4 for _ in range(4)]
        self.add_new_tile()
        self.add_new_tile()
        self.draw_board()

        self.root.bind("<Key>", self.key_handler)

    def draw_board(self):
        for row in range(4):
            for col in range(4):
                tile_value = self.board[row][col]
                label = tk.Label(self.root, text=str(tile_value) if tile_value != 0 else "",
                                 bg="lightgrey", font=("Arial", 24), width=4, height=2)
                label.grid(row=row, column=col, padx=5, pady=5)

    def key_handler(self, event):
        if event.keysym in ['Up', 'Down', 'Left', 'Right']:
            if self.can_move(event.keysym):
                self.move(event.keysym)
                self.add_new_tile()
                self.draw_board()
                if not self.can_move_any():
                    self.game_over()

    def move(self, direction):
        if direction == 'Up':
            self.move_up()
        elif direction == 'Down':
            self.move_down()
        elif direction == 'Left':
            self.move_left()
        elif direction == 'Right':
            self.move_right()

    def move_up(self):
        for col in range(4):
            column = self.get_column(col)
            merged = self.merge(column)
            self.set_column(col, merged)

    def move_down(self):
        for col in range(4):
            column = self.get_column(col)
            merged = self.reverse(self.merge(self.reverse(column)))
            self.set_column(col, merged)

    def move_left(self):
        for row in range(4):
            line = self.board[row]
            merged = self.merge(line)
            self.board[row] = merged

    def move_right(self):
        for row in range(4):
            line = self.board[row]
            merged = self.reverse(self.merge(self.reverse(line)))
            self.board[row] = merged

    def get_column(self, col):
        return [self.board[row][col] for row in range(4)]

    def set_column(self, col, column):
        for row in range(4):
            self.board[row][col] = column[row]

    def reverse(self, line):
        return line[::-1]

    def merge(self, line):
        non_zero = [num for num in line if num != 0]
        merged = []
        skip = False
        for i in range(len(non_zero)):
            if skip:
                skip = False
                continue
            if i + 1 < len(non_zero) and non_zero[i] == non_zero[i + 1]:
                merged.append(non_zero[i] * 2)
                skip = True
            else:
                merged.append(non_zero[i])
        return merged + [0] * (4 - len(merged))

    def add_new_tile(self):
        empty_spaces = [(r, c) for r in range(4) for c in range(4) if self.board[r][c] == 0]
        if empty_spaces:
            row, col = random.choice(empty_spaces)
            self.board[row][col] = 2

    def can_move(self, direction):
        test_board = [row[:] for row in self.board]
        self.move(direction)
        can_move = test_board != self.board
        self.board = test_board
        return can_move

    def can_move_any(self):
        for direction in ['Up', 'Down', 'Left', 'Right']:
            if self.can_move(direction):
                return True
        return False

    def game_over(self):
        game_over_label = tk.Label(self.root, text="游戏结束", bg="red", font=("Arial", 24), width=16, height=4)
        game_over_label.grid(row=0, column=0, columnspan=4, rowspan=4)

if __name__ == "__main__":
    root = tk.Tk()
    game = Game2048(root)
    root.mainloop()
