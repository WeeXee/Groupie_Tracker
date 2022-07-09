var script = document.createElement('script')
script.src = 'https://maps.googleapis.com/maps/api/js?key=AIzaSyDFs-nFa6EuWJFqJlfutT1xfVftj4-ZDVI&callback=initMap&libraries=&v=weekly&channel=2'
script.defer = true

window.initMap = function(x, y) {
    // The location of Uluru
    const uluru = { lat: 89.23/*x*/, lng:103 /*y*/ };
    // The map, centered at Uluru
    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 4,
        center: uluru,
    });
    // The marker, positioned at Uluru
    new google.maps.Marker({
        position: uluru,
        map: map,
    });
};

document.head.appendChild(script)