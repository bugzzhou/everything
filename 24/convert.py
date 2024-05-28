file_path = './sample.txt'  # 替换为你的文件路径

def parse_file(file_path='./sample.txt'):
    result = []
    with open(file_path, 'r', encoding='utf-8') as file:
        for line in file:
            if '－－－' in line:
                key, value = line.split('－－－')
                key = key.strip()
                value = value.strip()
                result.append({key: value})
    return result

# parsed_data = parse_file(file_path)
# print(len(parsed_data))