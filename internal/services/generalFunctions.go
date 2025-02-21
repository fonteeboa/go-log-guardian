/**
 * @file generalFunctions.go
 * @brief Este arquivo contém funções gerais utilizadas no serviço.
 */

package services

import (
	"fmt"
	"os"

	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

/**
 * @brief Verifica o ambiente e retorna um valor booleano.
 *
 * Esta função verifica se a variável de ambiente "DATABASE_TYPE" está definida.
 *
 * @return bool Retorna verdadeiro se a variável de ambiente "DATABASE_TYPE" estiver definida, caso contrário, retorna falso.
 */
func CheckEnvironment() bool {
	insertDB := os.Getenv("DATABASE_TYPE")
	return insertDB != ""
}

/**
 * @brief Imprime a mensagem de erro fornecida se a prioridade estiver definida como LOG_DEBUG.
 *
 * @param priority O nível de prioridade da mensagem de erro (domain.Priority).
 * @param genericErrMsg A mensagem de erro genérica (string).
 * @param errMsg A mensagem de erro específica (string).
 */
func Debug(priority domain.Priority, genericErrMsg string, errMsg string) {
	if priority == domain.LOG_DEBUG {
		fmt.Println(domain.PriorityToString[priority] + ": " + genericErrMsg + " " + errMsg)
	}
}
