

document.getElementById('addLinha').addEventListener('click', function() {
    const tabela = document.getElementById('tabelaProdutos').getElementsByTagName('tbody')[0];

    // Criar nova linha
    const novaLinha = document.createElement('tr');

    // Primeira célula: botão de remoção + campo editável para cores
    const novaCor = document.createElement('td');
    const btnRemove = document.createElement('button');
    btnRemove.className = 'delete-btn';
    btnRemove.innerText = 'X';
    btnRemove.addEventListener('click', function() {
        novaLinha.remove();
    });

    const spanEditavel = document.createElement('span');
    spanEditavel.contentEditable = 'true';
    spanEditavel.innerText = 'Cor';

    novaLinha.appendChild(btnRemove);
    novaCor.appendChild(spanEditavel);
    novaLinha.appendChild(novaCor);

    // Demais células: caixas de seleção para os tamanhos
    const tamanhos = ['00', '01', '02', '04', '06', '08', '10'];
    tamanhos.forEach(tamanho => {
        const novaCelula = document.createElement('td');
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.id = `nova_${tamanho}`;
        novaCelula.appendChild(checkbox);
        novaLinha.appendChild(novaCelula);
    });

    // Adicionar a nova linha à tabela
    tabela.appendChild(novaLinha);
});

// Função para alterar os cabeçalhos com base na seleção da categoria
document.getElementById('categoria').addEventListener('change', function() {
    const tabela = document.getElementById('tabelaProdutos');
    const colunasBaby = ['BB1', 'BB2', 'BB3', 'BB4', '00', '01', '02'];
    const colunasInfantil = ['04', '06', '08', '10', '12', '14', '16'];
    const colunasAdulto = ['PP', 'P', 'M', 'G', 'GG', 'EXG', 'Especial'];

    let novasColunas = [];

    switch (this.value) {
        case 'baby':
            novasColunas = colunasBaby;
            break;
        case 'infantil':
            novasColunas = colunasInfantil;
            break;
        case 'adulto':
            novasColunas = colunasAdulto;
            break;
    }

    // Atualizar os cabeçalhos da tabela
    const headers = tabela.getElementsByTagName('th');
    for (let i = 0; i < novasColunas.length; i++) {
        headers[i + 2].innerText = novasColunas[i];
    }
});

// Executar a seleção de adulto por padrão ao carregar a página
document.addEventListener('DOMContentLoaded', function() {
    const selectCategoria = document.getElementById('categoria');
    selectCategoria.dispatchEvent(new Event('change'));
});
