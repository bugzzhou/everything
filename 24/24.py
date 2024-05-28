import tkinter as tk
import random
import re
import convert

number_combinations = convert.parse_file()

class Game24:
    def __init__(self, root):
        self.root = root
        self.root.title("24点游戏")
        
        # 设置默认窗口大小
        self.root.geometry("800x600")

        # 创建显示数字的标签
        self.label_numbers = tk.Label(root)
        self.label_numbers.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        # 创建输入框和确定按钮
        self.entry = tk.Entry(root, state='readonly')
        self.entry.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)
        self.button_submit = tk.Button(root, text="确定", command=self.calculate_result)
        self.button_submit.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        # 创建显示结果的标签
        self.label_result = tk.Label(root, text="")
        self.label_result.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        # 创建提示按钮
        self.button_hint = tk.Button(root, text="提示", command=self.show_hint)
        self.button_hint.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        # 创建随机按钮
        self.button_random = tk.Button(root, text="随机", command=self.random_combination)
        self.button_random.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        # 创建输入键盘
        self.buttons_frame = tk.Frame(root)
        self.buttons_frame.pack(fill=tk.BOTH, expand=True, padx=10, pady=10)

        self.buttons = []
        button_texts = [
            '1', '2', '3', '+',
            '4', '5', '6', '-',
            '7', '8', '9', '*',
            '(', '0', '10', ')', '/',
            '退格', '清除'
        ]

        for row in range(5):
            self.buttons_frame.grid_rowconfigure(row, weight=1)
        for col in range(4):
            self.buttons_frame.grid_columnconfigure(col, weight=1)

        for idx, text in enumerate(button_texts):
            if text == '退格':
                button = tk.Button(self.buttons_frame, text=text, command=self.delete_last)
            elif text == '清除':
                button = tk.Button(self.buttons_frame, text=text, command=self.clear_entry)
            else:
                button = tk.Button(self.buttons_frame, text=text, command=lambda t=text: self.append_to_entry(t))
            button.grid(row=idx//4, column=idx%4, sticky=tk.NSEW, padx=5, pady=5)
            self.buttons.append(button)

        # 初始化组合
        self.random_combination()

    def random_combination(self):
        combination = random.choice(number_combinations)
        self.current_combination = list(combination.keys())[0]
        self.current_value = combination[self.current_combination]
        self.label_numbers.config(text="数字：" + self.current_combination)
        self.update_keyboard(self.current_combination)
        self.clear_entry()

    def calculate_result(self):
        expression = self.entry.get()
        try:
            # 提取输入的数字并排序
            input_numbers = sorted(re.findall(r'\d+', expression), key=int)
            # 提取随机选择的数字组合并排序
            selected_numbers = sorted(self.label_numbers.cget("text").split("：")[1].split(","), key=int)

            if input_numbers != selected_numbers:
                self.label_result.config(text="输入字符不太对")
                print(input_numbers)
                print(selected_numbers)
                return

            result = eval(expression)
            if result == 24:
                self.label_result.config(text="成功")
                self.random_combination()
                self.clear_entry()
            else:
                self.label_result.config(text="{} = {}".format(expression, result))
        except Exception as e:
            self.label_result.config(text="表达式无效: " + str(e))

    def update_keyboard(self, combination):
        numbers = combination.split(",")
        for button in self.buttons:
            if button["text"] in numbers or not button["text"].isdigit():
                button.config(state=tk.NORMAL)
            else:
                button.config(state=tk.DISABLED)

    def append_to_entry(self, value):
        self.entry.config(state=tk.NORMAL)
        self.entry.insert(tk.END, value)
        self.entry.config(state='readonly')

    def clear_entry(self):
        self.entry.config(state=tk.NORMAL)
        self.entry.delete(0, tk.END)
        self.entry.config(state='readonly')
        self.label_result.config(text="")

    def delete_last(self):
        self.entry.config(state=tk.NORMAL)
        self.entry.delete(len(self.entry.get())-1, tk.END)
        self.entry.config(state='readonly')

    def show_hint(self):
        self.label_result.config(text="直接显示结果：{}".format(self.current_value))

# 创建主窗口并运行
if __name__ == "__main__":
    root = tk.Tk()
    game = Game24(root)
    root.mainloop()
