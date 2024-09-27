// Exemplo de array de clientes vindos de uma API (dados do banco)
const clients = [
    { nome: 'João Silva', cidade: 'São Paulo', estado: 'SP' },
    { nome: 'Ana Souza', cidade: 'Rio de Janeiro', estado: 'RJ' },
    { nome: 'Pedro Costa', cidade: 'Belo Horizonte', estado: 'MG' },
    { nome: 'Maria Oliveira', cidade: 'Curitiba', estado: 'PR' },
    { nome: 'Lucas Martins', cidade: 'Porto Alegre', estado: 'RS' },
    { nome: 'Fernanda Lima', cidade: 'Fortaleza', estado: 'CE' },
    { nome: 'Carlos Pereira', cidade: 'Salvador', estado: 'BA' },
    { nome: 'Juliana Ferreira', cidade: 'Brasília', estado: 'DF' },
    { nome: 'Paulo Henrique', cidade: 'Goiânia', estado: 'GO' },
    { nome: 'Roberta Menezes', cidade: 'Recife', estado: 'PE' },
    { nome: 'Rodrigo Azevedo', cidade: 'Manaus', estado: 'AM' },
    { nome: 'Camila Rodrigues', cidade: 'Natal', estado: 'RN' },
    { nome: 'Vinícius Almeida', cidade: 'Florianópolis', estado: 'SC' },
    { nome: 'Tatiana Silva', cidade: 'Belém', estado: 'PA' },
    { nome: 'Marcos Rocha', cidade: 'São Luís', estado: 'MA' }
];

const tbody = document.getElementById('client-table-body');
const searchInput = document.getElementById('search-input');
const nomeField = document.getElementById('nome-cliente'); // Campo de formulário para o nome
const cidadeField = document.getElementById('cidade-cliente'); // Campo de formulário para o cidade
const estadoField = document.getElementById('estado-cliente'); // Campo de formulário para o estado

// Função para preencher a tabela com os dados dos clientes
// function renderTable(data) {
//     const attBtn = document.getElementById('atualizar_btn');
//     tbody.innerHTML = ''; // Limpa o corpo da tabela
//     data.forEach((client, index) => {
//         const row = document.createElement('tr');
//         row.innerHTML = `
//             <td>${client.nome}</td>
//             <td>${client.cidade}</td>
//             <td>${client.estado}</td>
//         `;
//         row.addEventListener('click', function() {
//             // Quando a linha for clicada, preencher o campo de formulário com o nome do cliente
//             nomeField.value = client.nome;
//             cidadeField.value = client.cidade;
//             estadoField.value = client.estado;

//             attBtn.textContent = 'Atualizar';
//         });
//         tbody.appendChild(row);
//     });
// }

function handleClientsTableClick() {
    const tabela = document.getElementById('clients-table');

    tabela.addEventListener('click', (event) => {
        const linha = event.target.parentElement; // Obtém a linha da tabela clicada
        if (linha && linha.cells) {
        // Obtém os dados das células da linha
        const nome = linha.cells[0].innerText;
        const cidade = linha.cells[1].innerText;
        const estado = linha.cells[2].innerText;

        nomeField.value = nome
        cidadeField.value = cidade
        estadoField.value = estado

        }
    });
}

handleClientsTableClick();

// Função para filtrar clientes com base na pesquisa
function filterClients() {
    const query = searchInput.value.toLowerCase();
    const filteredClients = clients.filter(client => 
        client.nome.toLowerCase().includes(query) ||
        client.cidade.toLowerCase().includes(query) ||
        client.estado.toLowerCase().includes(query)
    );
    renderTable(filteredClients);
    

}

// Inicializa a tabela com todos os clientes
// renderTable(clients);

// Adiciona um ouvinte de evento ao campo de pesquisa
searchInput.addEventListener('input', filterClients);


window.onload = function() {
    clicar(); // Executa a função clicar quando a página terminar de carregar
};

function clicar() {
    const attBtn = document.getElementById('atualizar_btn');
    const novoClienteBtn = document.getElementById('novo_cliente_btn');
    const formulario = document.querySelector('form');

    // Adiciona o event listener ao botão
    novoClienteBtn.addEventListener('click', function() {
        // Altera o texto do botão
        attBtn.textContent = 'Cadastrar';  // Muda o texto para 'Atualizando...'
        formulario.querySelectorAll('input[type="text"]').forEach(input => {
            input.value = '';
        });
    });
}

// Chama a função imediatamente após sua definição
clicar();


const formContainerSelect = document.getElementById('form-container-select');
const createClientForm = document.getElementById('novo_cliente');
const updateClientForm = document.getElementById('atualizar_cliente');

formContainerSelect.onchange = (event) => {
    selectedOption = event.target.selectedIndex

    switch (selectedOption) {
        case 0: // Cadastrar cliente
            createClientForm.style.display = "block";
            updateClientForm.style.display = "none";
        break;
        case 1: // Atualizar cliente
            createClientForm.style.display = "none";
            updateClientForm.style.display = "block";
        break;
    }
}