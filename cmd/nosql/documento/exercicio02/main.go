package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dao, err := NewPessoaDocDAOImpl("")
	if err != nil {
		fmt.Println("Falha ao inicializar DAO:", err)
		return
	}
	defer func() {
		if err := dao.Close(context.Background()); err != nil {
			fmt.Println("Erro ao fechar ligação:", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menu:")
		fmt.Println("1 - Create")
		fmt.Println("2 - Read")
		fmt.Println("3 - Update")
		fmt.Println("4 - Delete")
		fmt.Println("0 - Fim")

		opcao, err := readInt(reader, "Opção: ")
		if err != nil {
			fmt.Println("Entrada inválida, introduz um número.")
			continue
		}

		switch opcao {
		case 1:
			createPessoa(reader, dao)
		case 2:
			readPessoaByID(reader, dao)
		case 3:
			updatePessoa(reader, dao)
		case 4:
			deletePessoa(reader, dao)
		case 0:
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
	text, err := readLine(reader, prompt)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(text)
}

func promptPessoa(reader *bufio.Reader) (Pessoa, error) {
	var p Pessoa
	var err error

	if p.ID, err = readInt(reader, "id: "); err != nil {
		return p, err
	}
	if p.Nome, err = readLine(reader, "Nome: "); err != nil {
		return p, err
	}
	if p.Sobrenome, err = readLine(reader, "Sobrenome: "); err != nil {
		return p, err
	}
	if p.Telefone, err = readLine(reader, "Telefone: "); err != nil {
		return p, err
	}
	if p.Idade, err = readInt(reader, "Idade: "); err != nil {
		return p, err
	}

	return p, nil
}

func createPessoa(reader *bufio.Reader, dao PessoaDAO) {
	fmt.Println("Create:")
	p, err := promptPessoa(reader)
	if err != nil {
		fmt.Println("Erro nos dados introduzidos:", err)
		return
	}
	if err := dao.Create(p); err != nil {
		fmt.Println(err)
	}
}

func updatePessoa(reader *bufio.Reader, dao PessoaDAO) {
	fmt.Println("Update:")
	p, err := promptPessoa(reader)
	if err != nil {
		fmt.Println("Erro nos dados introduzidos:", err)
		return
	}
	if err := dao.Update(p); err != nil {
		fmt.Println(err)
	}
}

func readPessoaByID(reader *bufio.Reader, dao PessoaDAO) {
	fmt.Println("Read:")
	id, err := readInt(reader, "Id: ")
	if err != nil {
		fmt.Println("Id inválido.")
		return
	}
	p, err := dao.Read(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if p == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(p)
}

func deletePessoa(reader *bufio.Reader, dao PessoaDAO) {
	fmt.Println("Delete:")
	id, err := readInt(reader, "Id: ")
	if err != nil {
		fmt.Println("Id inválido.")
		return
	}
	if err := dao.Delete(id); err != nil {
		fmt.Println(err)
	}
}
