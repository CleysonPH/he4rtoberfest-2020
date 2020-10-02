import sys


def hello_world(name: str = "He4rtoberfest") -> str:
    return f"Helo {name}"


if __name__ == "__man__":
    nome = sys.argv[0]
    output = hello_world(nome)
    print(output)
