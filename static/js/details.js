// static/js/details.js
$(document).ready(function () {
  const urlParams = new URLSearchParams(window.location.search);
  const propertyId = urlParams.get("id");
  if (propertyId) {
    loadPropertyDetails(propertyId);
  }
});

function loadPropertyDetails(propertyId) {
  $.ajax({
    url: `/v1/property/details?property_id=${propertyId}`,
    method: "GET",
    success: function (response) {
      if (response.success) {
        displayPropertyDetails(response.property);
      }
    },
    error: function (error) {
      console.error("Error loading property details:", error);
    },
  });
}

function displayPropertyDetails(property) {
  // Update breadcrumbs with clickable links
  document.getElementById("breadcrumbs").innerHTML = property.breadcrumbs
    .map(
      (crumb, index) => `
        <span class="cursor-pointer hover:text-blue-600" onclick="handleBreadcrumbClick('${crumb}', ${index})">
            ${crumb}
        </span>
        ${index < property.breadcrumbs.length - 1 ? " > " : ""}
    `
    )
    .join("");

  let roomsInfo = "";
  let bedsInfo = "";
  // Conditionally add bedrooms info
  if (property.bedrooms > 0) {
    roomsInfo += `<span>${property.bedrooms} Bedrooms </span>`;
    bedsInfo += `<span>${property.bedrooms} Bedrooms</span>`;
  }

  // Conditionally add bathrooms info
  if (property.bathrooms > 0) {
    roomsInfo += `<span>${property.bathrooms} Bathrooms </span>`;
  }

  // Conditionally add guest count info
  if (property.guest_count > 0) {
    roomsInfo += `<span>${property.guest_count} Guests </span>`;
  }

  const detailsHTML = `
        <div class="grid grid-cols-1 gap-8">
            <div>
                <div class="md:grid md:grid-cols-2 md:gap-2">
                    <img src="${
                      property.images[0] || "/static/images/placeholder.jpg"
                    }" alt="${
    property.hotel_name
  }" class="md:w-full w-full h-96 object-cover rounded-lg">
                    <div class="grid grid-cols-2 gap-2 mt-2 md:mt-0">
                        ${property.images
                          .slice(1, 5)
                          .map(
                            (img) => `
                            <img src="${img}" alt="" class="w-full h-48 object-cover rounded-lg cursor-pointer"
                                onclick="updateMainImage(this.src)">
                        `
                          )
                          .join("")}
                    </div>
                </div>
            </div>
            <div>
                <h1 class="text-3xl font-bold text-blue-900 mb-4">${bedsInfo} ${
    property.hotel_name
  } ${property.city_in_trans}</h1>
                <div class="md:items-center mb-4">
                    <span class="text-blue-600 font-bold ">${
                      property.rating
                    }</span>
                    <span class="md:ml-2 text-gray-600">(${
                      property.review_count
                    } Reviews)</span>
                    <span class="md:ml-2 text-gray-600">${roomsInfo}</span>
                    <span class="md:ml-2 text-gray-600">${property.amenities
                      .map(
                        (amenity) => `
                        <span class="text-gray-600 ml-2">${amenity}</span>
                    `
                      )
                      .join("")}</span>
                    
                </div>
                <div class="mb-6">
                    <h2 class="text-xl font-bold mb-2">Description</h2>
                    <p class="text-gray-600">${
                      property.description || "No description available"
                    }</p>
                </div>
                <div class="mb-6">
                    <h2 class="text-xl font-bold mb-2">Amenities</h2>
                    <div class="grid grid-cols-2 gap-2">
                        ${property.amenities
                          .map(
                            (amenity) => `
                            <div class="text-gray-600">${amenity}</div>
                        `
                          )
                          .join("")}
                    </div>
                </div>
                <button class="w-full bg-emerald-500 text-white py-3 rounded-lg text-lg font-bold">
                    Check Availability
                </button>
            </div>
        </div>
    `;

  document.getElementById("propertyDetails").innerHTML = detailsHTML;
}
function handleBreadcrumbClick(locationName, level) {
  window.location.href = `/?location=${encodeURIComponent(locationName)}`;
}

function updateMainImage(src) {
  const mainImage = document.querySelector(".h-96");
  if (mainImage) {
    mainImage.src = src;
  }
}
