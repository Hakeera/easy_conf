// Exemplo de array de clientes vindos de uma API (dados do banco)
const clientes = [
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

const tbody = document.getElementById('cliente-table-body');
const searchInput = document.getElementById('search-input');
const nomeField = document.getElementById('nome'); // Campo de formulário para o nome
const cidadeField = document.getElementById('cidade'); // Campo de formulário para o cidade
const estadoField = document.getElementById('estado'); // Campo de formulário para o estado

// Função para preencher a tabela com os dados dos clientes
function renderTable(data) {
    const attBtn = document.getElementById('atualizar_btn');
    tbody.innerHTML = ''; // Limpa o corpo da tabela
    data.forEach((client, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${client.nome}</td>
            <td>${client.cidade}</td>
            <td>${client.estado}</td>
        `;
        row.addEventListener('click', function() {
            // Quando a linha for clicada, preencher o campo de formulário com o nome do cliente
            nomeField.value = client.nome;
            cidadeField.value = client.cidade;
            estadoField.value = client.estado;

            attBtn.textContent = 'Atualizar';
        });
        tbody.appendChild(row);
    });
}

// Função para filtrar clientes com base na pesquisa
function filterClientes() {
    const query = searchInput.value.toLowerCase();
    const filteredClientes = clientes.filter(client => 
        client.nome.toLowerCase().includes(query) ||
        client.cidade.toLowerCase().includes(query) ||
        client.estado.toLowerCase().includes(query)
    );
    renderTable(filteredClientes);
    

}

// Inicializa a tabela com todos os clientes
renderTable(clientes);

// Adiciona um ouvinte de evento ao campo de pesquisa
searchInput.addEventListener('input', filterClientes);


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