let currentLocationId = 1; // Default location ID

// Load properties on page load
$(document).ready(function() {
    loadPropertiesByLocation(currentLocationId);
    
    // Setup search input event listener
    $('#locationSearch').on('input', function() {
        const searchQuery = $(this).val().toLowerCase();
        if (searchQuery.length > 2) {
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

function displayProperties(properties) {
    const propertyGrid = document.getElementById('propertyGrid');
    propertyGrid.innerHTML = '';

    properties.forEach(property => {
        const propertyCard = `
            <div class="bg-white rounded-lg shadow-md overflow-hidden">
                <div class="relative">
                    <img src="${property.images[0] || '/static/images/placeholder.jpg'}" alt="${property.hotel_name}" class="w-full h-64 object-cover">
                    <div class="absolute top-4 right-4 flex space-x-2">
                        <button class="p-2 bg-white rounded-full shadow-md">
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" />
                            </svg>
                        </button>
                    </div>
                </div>
                <div class="p-4">
                    <div class="flex items-center mb-2">
                        <span class="text-blue-600 font-bold">${property.rating}</span>
                        <span class="ml-2 text-gray-600">(${property.review_count} Reviews)</span>
                    </div>
                    <h3 class="text-xl font-bold mb-2">
                        <a href="/property-details?id=${property.property_id}" target="_blank" class="text-blue-900 hover:underline">
                            ${property.hotel_name}
                        </a>
                    </h3>
                    <div class="text-gray-600 mb-2">
                        ${property.breadcrumbs.map((crumb, index) => `
                            <span class="cursor-pointer hover:text-blue-600" onclick="handleBreadcrumbClick('${crumb}', ${index})">${crumb}</span>
                            ${index < property.breadcrumbs.length - 1 ? ' > ' : ''}
                        `).join('')}
                    </div>
                    <div class="flex items-center justify-between mt-4">
                        <div class="text-lg font-bold">From ${property.price}</div>
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