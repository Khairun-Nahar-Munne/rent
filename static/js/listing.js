let currentLocationId = 1; // Default location ID

// Load properties on page load
$(document).ready(function() {
    loadPropertiesByLocation(currentLocationId);
    
    // Setup search input event listener
    $('#locationSearch').on('input', function() {
        const searchQuery = $(this).val().toLowerCase();
        if (searchQuery.length > 1) {
            searchLocations(searchQuery);
        }
    });
});

function loadPropertiesByLocation(locationId) {
    $.ajax({
        url: '/v1/property/list',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const location = response.locations.find(loc => loc.id === locationId);
                if (location) {
                    updateBreadcrumb(location);
                    updatePageTitle(location);
                    displayProperties(location.properties);
                }
            }
        },
        error: function(error) {
            console.error('Error loading properties:', error);
        }
    });
}

function updateBreadcrumb(location) {
    const breadcrumbHTML = `
        <span class="cursor-pointer hover:text-blue-600" onclick="loadPropertiesByLocation(${location.id})">
            ${location.value}
        </span>
    `;
    document.getElementById('locationBreadcrumb').innerHTML = breadcrumbHTML;
}

function updatePageTitle(location) {
    const pageTitle = document.getElementById('pageTitle');
    if (location.properties.length > 0 && location.properties[0].breadcrumbs.length > 0) {
        pageTitle.textContent = `Rent By Owner ${location.properties[0].breadcrumbs[0]} - Vacation Rentals in ${location.properties[0].breadcrumbs[0]}`;
    } else {
        pageTitle.textContent = `Rent By Owner - Vacation Rentals`;
    }
}

function displayProperties(properties) {
    const propertyGrid = document.getElementById('propertyGrid');
    propertyGrid.innerHTML = '';

    properties.forEach(property => {
        let roomsInfo = '';

        // Conditionally add bedrooms info
        if (property.bedrooms > 0) {
            roomsInfo += `<span>${property.bedrooms} Bedrooms</span>`;
        }

        // Conditionally add bathrooms info
        if (property.bathrooms > 0) {
            roomsInfo += `<span>${property.bathrooms} Bathrooms</span>`;
        }

        // Conditionally add guest count info
        if (property.guest_count > 0) {
            roomsInfo += `<span>${property.guest_count} Guests</span>`;
        }

        const propertyCard = `
            <div class="bg-white rounded-lg shadow-md overflow-hidden">
                <div class="relative">
                    <img src="${property.images[0] || '/static/images/placeholder.jpg'}" alt="${property.hotel_name}" class="w-full h-64 object-cover">
                    <div class="absolute  bottom-4 left-4 bg-white text-blue-900 font-bold px-2 py-1 rounded-lg shadow-md">
                        From ${property.price}
                    </div>
                </div>
                <div class="p-4">
                    <div class="flex items-center justify-between mb-2 text-sm">
                        <div> <span class="text-blue-600 font-bold text-sm">${property.rating}</span>
                    <span class="ml-2 text-gray-600">(${property.review_count} Reviews)</span></div>
                    <span class="ml-2 text-gray-600">${property.type}</span>
                    </div>
                    <h3 class="text-xl font-bold mb-2">
                        <a href="/property-details?id=${property.property_id}" target="_blank" class="text-blue-900 hover:underline">
                            ${property.hotel_name}
                        </a>
                    </h3>
                    <div class="text-gray-600 text-sm">${property.amenities.join('  *')}</div>
                    <div class="text-gray-600 mb-2 text-sm">
                        ${property.breadcrumbs.map((crumb, index) => `
                            <span class="cursor-pointer text-md hover:text-blue-600 text-sm" onclick="handleBreadcrumbClick('${crumb}', ${index})">${crumb}</span>
                            ${index < property.breadcrumbs.length - 1 ? ' > ' : ''}
                        `).join('')}
                    </div>
                    <div class="flex items-center justify-between mt-4">

                        <div class="bg-blue-900 text-white font-bold px-2 py-1 rounded-lg shadow-md">
                            Booking.com
                        </div>
                        <button onclick="window.open('/property-details?id=${property.property_id}', '_blank')" 
                                class="bg-emerald-500 text-white px-4 py-2 rounded-lg">
                            View Availability
                        </button>
                    </div>
                </div>
            </div>
        `;
        propertyGrid.innerHTML += propertyCard;
    });
}

function handleBreadcrumbClick(locationName, level) {
    $.ajax({
        url: '/v1/property/list',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const location = response.locations.find(loc => 
                    loc.value.toLowerCase() === locationName.toLowerCase()
                );
                if (location) {
                    currentLocationId = location.id;
                    loadPropertiesByLocation(location.id);
                }
            }
        }
    });
}

function searchLocations(query) {
    $.ajax({
        url: '/v1/property/list',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const matchedLocation = response.locations.find(loc => 
                    loc.value.toLowerCase().includes(query.toLowerCase())
                );
                if (matchedLocation) {
                    currentLocationId = matchedLocation.id;
                    loadPropertiesByLocation(matchedLocation.id);
                }
            }
        }
    });
}