// static/js/details.js
$(document).ready(function() {
    const urlParams = new URLSearchParams(window.location.search);
    const propertyId = urlParams.get('id');
    if (propertyId) {
        loadPropertyDetails(propertyId);
    }
});

function loadPropertyDetails(propertyId) {
    $.ajax({
        url: `/v1/property/details?property_id=${propertyId}`,
        method: 'GET',
        success: function(response) {
            if (response.success) {
                displayPropertyDetails(response.property);
            }
        },
        error: function(error) {
            console.error('Error loading property details:', error);
        }
    });
}

function displayPropertyDetails(property) {
    // Update breadcrumbs with clickable links
    document.getElementById('breadcrumbs').innerHTML = property.breadcrumbs.map((crumb, index) => `
        <span class="cursor-pointer hover:text-blue-600" onclick="handleBreadcrumbClick('${crumb}', ${index})">
            ${crumb}
        </span>
        ${index < property.breadcrumbs.length - 1 ? ' > ' : ''}
    `).join('');

    const detailsHTML = `
        <div class="grid grid-cols-1 gap-8">
            <div>
                <div class="relative">
                    <img src="${property.images[0] || '/static/images/placeholder.jpg'}" alt="${property.hotel_name}" class="w-full h-96 object-cover rounded-lg">
                    <div class="grid grid-cols-4 gap-2 mt-2">
                        ${property.images.slice(1, 5).map(img => `
                            <img src="${img}" alt="" class="w-full h-24 object-cover rounded-lg cursor-pointer"
                                 onclick="updateMainImage(this.src)">
                        `).join('')}
                    </div>
                </div>
            </div>
            <div>
                <h1 class="text-3xl font-bold text-blue-900 mb-4">${property.hotel_name}</h1>
                <div class="flex items-center mb-4">
                    <span class="text-blue-600 font-bold text-xl">${property.rating}</span>
                    <span class="ml-2 text-gray-600">(${property.review_count} Reviews)</span>
                </div>
                <div class="grid grid-cols-2 gap-4 mb-6">
                    <div class="text-gray-600">
                        <div>Bedrooms: ${property.bedrooms}</div>
                        <div>Bathrooms: ${property.bathrooms}</div>
                        <div>Guests: ${property.guest_count}</div>
                    </div>
                    <div class="text-gray-600">
                        <div>Type: ${property.type}</div>
                        <div>Price: ${property.price}</div>
                    </div>
                </div>
                <div class="mb-6">
                    <h2 class="text-xl font-bold mb-2">Description</h2>
                    <p class="text-gray-600">${property.description || 'No description available'}</p>
                </div>
                <div class="mb-6">
                    <h2 class="text-xl font-bold mb-2">Amenities</h2>
                    <div class="grid grid-cols-2 gap-2">
                        ${property.amenities.map(amenity => `
                            <div class="text-gray-600">${amenity}</div>
                        `).join('')}
                    </div>
                </div>
                <button class="w-full bg-emerald-500 text-white py-3 rounded-lg text-lg font-bold">
                    Check Availability
                </button>
            </div>
        </div>
    `;

    document.getElementById('propertyDetails').innerHTML = detailsHTML;
}

function handleBreadcrumbClick(locationName, level) {
    window.location.href = `/?location=${encodeURIComponent(locationName)}`;
}

function updateMainImage(src) {
    const mainImage = document.querySelector('.h-96');
    if (mainImage) {
        mainImage.src = src;
    }
}