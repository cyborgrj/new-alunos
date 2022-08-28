package custom_errors

import "errors"

var (
	ErrCPFInvalido      = errors.New("CPF inválido")
	ErrIDInvalido       = errors.New("ID de aluno inválido")
	ErrCPFJaCadastrado  = errors.New("CPF já cadastrado")
	ErrCEPnaoRecuperado = errors.New("erro ao recuperar o endereço a partir do cep")
	ErrCEPnaoInformado  = errors.New("CEP não informado")
	ErrCEPInvalido      = errors.New("CEP Inválido")
	ErrAlunoInvalido    = errors.New("Os dados inseridos para o aluno são inválidos. Verificar.")
)
