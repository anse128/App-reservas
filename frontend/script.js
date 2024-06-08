document.getElementById('auth-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Simulación de autenticación
    if (username && password) {
        document.getElementById('auth-section').style.display = 'none';
        document.getElementById('reserve-section').style.display = 'block';
    } else {
        alert('Por favor, ingrese un nombre de usuario y contraseña válidos');
    }
});

document.getElementById('search-button').addEventListener('click', function() {
    const query = document.getElementById('search-query').value;

    // Simulación de búsqueda de autos
    const carList = document.getElementById('car-list');
    carList.innerHTML = '';

    const cars = [
        { id: 1, marca: 'Toyota', modelo: 'Corolla', combustible: 'Gasolina', imagen: 'images/corolla.jpg' },
        { id: 2, marca: 'Honda', modelo: 'Civic', combustible: 'Gasolina', imagen: 'images/civic.jpg' }
    ];

    cars.forEach(car => {
        const carCard = document.createElement('div');
        carCard.className = 'car-card';
        carCard.innerHTML = `
            <img src="${car.imagen}" alt="${car.modelo}">
            <h3>${car.marca} ${car.modelo}</h3>
            <p>Combustible: ${car.combustible}</p>
            <button onclick="reservarAuto(${car.id})">Reservar</button>
        `;
        carList.appendChild(carCard);
    });
});

function reservarAuto(carId) {
    // Simulación de reserva de auto
    alert(`Auto con ID ${carId} reservado!`);
}
