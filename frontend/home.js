// Seleciona a div itsel e o botão
const itselDiv = document.getElementById('itsel');
const toggleButton = document.getElementById('toggle-arrow');

// Função para alternar a visibilidade da div
toggleButton.addEventListener('click', function() {
    if (itselDiv.classList.contains('visible')) {
        // Oculta a div e muda a seta para baixo
        itselDiv.classList.remove('visible');
        toggleButton.innerHTML = '&#x25BC;'; // Seta para baixo
    } else {
        // Mostra a div e muda a seta para cima
        itselDiv.classList.add('visible');
        toggleButton.innerHTML = '&#x25B2;'; // Seta para cima
    }
});