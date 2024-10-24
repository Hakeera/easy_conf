// main.go: Inicializa o servidor HTTP e gerencia o shutdown gracioso

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"ezConf/internal/server"
)

// gracefulShutdown realiza o shutdown gracioso do servidor HTTP
// Ele aguarda sinais de interrupção (SIGINT ou SIGTERM), notifica o servidor
// para finalizar as requisições atuais e dá um prazo de 5 segundos para desligar.
// Parâmetros:
// - apiServer: Instância do servidor HTTP.
// - done: Canal para sinalizar quando o shutdown for concluído.
func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Cria um contexto que escuta sinais de interrupção do SO (SIGINT, SIGTERM).
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Aguarda o sinal de interrupção.
	<-ctx.Done()

	// Loga que o servidor iniciou o processo de shutdown gracioso.
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// Define um timeout de 5 segundos para o servidor finalizar as requisições pendentes.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Inicia o processo de shutdown do servidor com o timeout definido.
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	// Loga o encerramento do servidor.
	log.Println("Server exiting")

	// Notifica a goroutine principal que o shutdown foi concluído.
	done <- true
}

// main inicializa o servidor HTTP e gerencia o ciclo de vida dele.
// Cria um canal para controlar o shutdown gracioso, inicializa o servidor, 
// e espera até que o processo de shutdown seja finalizado.
func main() {
	// Inicializa o servidor HTTP através da função NewServer no pacote `server`.
	server := server.NewServer()

	// Cria um canal para sinalizar quando o shutdown for finalizado.
	done := make(chan bool, 1)

	// Executa a função de shutdown gracioso em uma goroutine separada.
	go gracefulShutdown(server, done)

	// Inicia o servidor HTTP e escuta as requisições.
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		// Em caso de erro crítico no servidor, exibe o erro e encerra o programa.
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Espera até que o processo de shutdown gracioso seja concluído.
	<-done
	log.Println("Graceful shutdown complete.")
}
