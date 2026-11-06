package main

import "fmt"

var dao PessoaDAO = NewPessoaDAOImpl()

func create() {
	var pessoa Pessoa
	fmt.Print("ID: ")
	fmt.Scanf("%s", &pessoa.ID)
	fmt.Print("Nome: ")
	fmt.Scanf("%s", &pessoa.Nome)
	fmt.Print("Telefone: ")
	fmt.Scanf("%s", &pessoa.Telefone)
	fmt.Print("Email: ")
	fmt.Scanf("%s", &pessoa.Email)

	err := dao.Create(pessoa)
	if err != nil {
		fmt.Println("Erro ao criar pessoa:", err)
	} else {
		fmt.Println("Pessoa criada com sucesso!")
	}
}

func read() {
	var id string
	fmt.Print("ID: ")
	fmt.Scanf("%s", &id)

	pessoa, err := dao.Read(id)
	if err != nil {
		fmt.Println("Erro ao ler pessoa:", err)
	} else {
		fmt.Printf("Pessoa encontrada: %+v\n", pessoa)
	}
}

func update() {
	create()
}

func del() {
	var id string
	fmt.Print("ID: ")
	fmt.Scanf("%s", &id)

	err := dao.Delete(id)
	if err != nil {
		fmt.Println("Erro ao deletar pessoa:", err)
	} else {
		fmt.Println("Pessoa deletada com sucesso!")
	}
}

func main() {
	var opcao int
	for ok := true; ok; ok = opcao != 0 {
		msg := `
                Menu:
                1 - Create
                2 - Read
                3 - Update
                4 - Delete
                0 - Fim
                `
		fmt.Println(msg)
		fmt.Print("Escolha uma opção: ")
		fmt.Scanf("%d", &opcao)
		switch opcao {
		case 1:
			create()
		case 2:
			read()
		case 3:
			update()
		case 4:
			del()
		}
	}
}
