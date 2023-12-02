document.addEventListener('DOMContentLoaded', function () {
    let addProductForm = document.getElementById('addProductForm');
    addProductForm.addEventListener('submit', function (event) {
        event.preventDefault();

        let productName = document.getElementById('productName').value;
        let productPrice = document.getElementById('productPrice').value;
        let productImage = document.getElementById('productImage').value;
        let productDescription = document.getElementById('productDescription').value;

        addToCart(productName, productPrice, productImage, productDescription);

    });

    function addToCart(name, price, image, description) {
        let newCard = document.createElement('div');
        newCard.classList.add('col');
        newCard.innerHTML = `
            <div class="card shadow-sm">
                <img src="${image}" class="bd-placeholder-img card-img-top" width="100%" height="225" role="img" aria-label="Placeholder: ${name}" preserveAspectRatio="xMidYMid slice" focusable="false">
                <div class="card-body">
                    <p class="card-text">${name}</p>
                    <p class="card-text">${description}</p>
                    <div class="d-flex justify-content-between align-items-center">
                        <div class="btn-group">
                            <button type="button" class="btn btn-sm btn-outline-secondary">Купить</button>
                            <button type="button" class="btn btn-sm btn-outline-secondary">Купить в рассрочку</button>
                        </div>
                        <small class="text-body-secondary">${price} рублей</small>
                    </div>
                </div>
            </div>
        `;

        let container = document.querySelector('.row.row-cols-1.row-cols-sm-2.row-cols-md-3.g-3');
        
        container.appendChild(newCard);
    }
});
