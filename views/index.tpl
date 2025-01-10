<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Property Listings</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script src="/static/js/listing.js" defer></script>
</head>
<body class="bg-gray-100">
    <!-- Navigation -->
    <nav class="bg-blue-900 p-4">
        <div class="container mx-auto md:flex md:items-center md:justify-between">
            <div class="text-white text-2xl font-bold">Rent by owner</div>
            <div class="md:flex md:space-x-4">
                <input type="text" id="locationSearch" placeholder="United States of America" class="p-2 rounded-lg w-64">
                <button onclick="searchProperties()" class="bg-emerald-500 text-white mt-2 md:mt-0 px-6 py-2 rounded-lg">Search</button>
            </div>
        </div>
    </nav>

    <!-- Main Content -->
    <div class="container mx-auto mt-8">
        <!-- Breadcrumb -->
        <div id="locationBreadcrumb" class="text-sm text-gray-600 mb-4"></div>
        
        <h1 id="pageTitle" class="text-3xl font-bold text-blue-900 mb-6">Rent By Owner England - Vacation Rentals in England</h1>
        

        <!-- Property Grid -->
        <div id="propertyGrid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <!-- Properties will be loaded here via AJAX -->
        </div>
    </div>
</body>
</html>