package event

var fakeStreet1 = []string{
	"Maple Avenue", "Oak Street", "Elm Road", "Cedar Lane", "Pine Court", "Willow Way", "Birch Street", "Spruce Drive", "Aspen Lane", "Juniper Avenue", "Chestnut Road", "Cypress Court", "Sycamore Lane", "Poplar Street", "Alder Drive", "Magnolia Avenue", "Walnut Lane", "Hemlock Road", "Hawthorn Court", "Cherry Street", "Redwood Drive", "Beech Avenue", "Mulberry Lane", "Mahogany Road", "Olive Court", "Ash Street", "Hickory Way", "Yew Lane", "Ginkgo Avenue", "Rowan Road", "Fir Court", "Palm Lane", "Banana Street", "Apple Drive", "Peach Avenue", "Lemon Lane", "Lime Road", "Orange Court", "Grape Street", "Plum Way", "Berry Avenue", "Raspberry Road", "Blueberry Court", "Strawberry Lane", "Watermelon Street", "Pineapple Drive", "Mango Avenue", "Kiwi Lane", "Papaya Road", "Avocado Court",
}

var fakeStreet2 = []string{
	"Grove", "Lane", "Boulevard", "Avenue", "Court", "Drive", "Place", "Road", "Circle", "Terrace", "Way", "Alley", "Square", "Path", "Parkway", "Lane", "Highway", "Walk", "Crescent", "Loop", "Mews", "Row", "Close", "Cresent", "Park", "Trace", "Crossroad", "Ridge", "Run", "Heights", "Gardens", "View", "Plaza", "Corner", "Passage", "Lane", "Alleyway", "Court", "Freeway", "Esplanade", "Chase", "Rise", "Bend", "Street", "Green", "Extension", "Way", "Square", "Parade", "Promenade",
}

var fakeCity = []string{
	"Aurora", "Savannah", "Kingston", "Valencia", "Adelaide", "Brasília", "Marseille", "Kyoto", "Wellington", "Vancouver", "Alexandria", "Budapest", "Nairobi", "Dublin", "Helsinki", "Casablanca", "Lima", "Perth", "Manila", "Oslo", "San Diego", "Istanbul", "Prague", "Buenos Aires", "Montreal", "Moscow", "Seoul", "Cairo", "Sydney", "Zurich",
}

var fakeState = []string{
	"California", "Texas", "Florida", "New York", "Pennsylvania", "Illinois", "Ohio", "Georgia", "Michigan", "North Carolina",
}

var fakePostalCode = []string{
	"90210", "77002", "EC1A 1BB", "SW1A 1AA", "M1 1AA", "WC2N 5DU", "W1A 0AX", "AB10 1AB", "G2 1AB", "EH1 1RH", "10001", "60601", "90292", "02108", "75201", "94102", "10017", "60606", "90015", "28202", "15222", "10005", "98101", "94103", "75202", "60622",
}

var fakeCountry = []string{
	"United States", "Canada", "Brazil", "France", "Germany", "India", "Japan", "Australia", "Italy", "South Africa", "Mexico", "Argentina", "China", "Spain", "United Kingdom", "Russia", "Egypt", "Nigeria", "Saudi Arabia", "Turkey", "Sweden", "Greece", "Thailand", "South Korea", "Malaysia", "Indonesia", "New Zealand", "Switzerland", "Pakistan", "Iran", "Colombia", "Chile", "Netherlands", "Belgium", "Peru", "Venezuela", "Denmark", "Finland", "Norway", "Portugal", "Ireland", "Vietnam", "Austria", "Philippines", "Hungary", "Czech Republic", "Poland", "Ukraine", "Israel", "Singapore",
}

var fakeCardNumber = []string{
	"4485 2466 8492 5953", "5241 3498 4671 1236", "3769 2890 4546 208", "6011 3758 5327 0138", "3548 2697 1099 3465", "4916 3472 9375 2041", "3792 2668 5404 981", "5412 0398 7615 2893", "6304 3871 2456 9772", "4024 1605 8735 2263", "3537 4290 8721 4767", "6373 5681 9423 5480", "5253 6089 1237 9442", "3019 7748 2659 8336", "4916 8830 5748 7107", "3621 1699 8725 3906", "5426 5281 3675 2971", "6761 2902 5403 7742", "4539 6072 8813 1609", "5141 5103 2489 5427", "3058 8786 9469 1808", "6382 0773 9208 0630", "3721 1546 9603 992", "4024 0071 2835 2816", "3000 2371 6271 7714", "3701 8842 2938 221", "4510 9285 3671 3730", "3761 3299 6811 165", "5421 8095 1024 3486", "6334 5687 2913 9277", "5140 9656 5342 9516", "3578 3986 9535 6788", "6011 6616 1872 9692", "3710 8634 9532 115",
}

var fakeCvv = []string{
	"123", "456", "789", "234", "567", "890", "345", "678", "901", "432", "765", "098", "321", "654", "987", "210", "543", "876", "109", "432", "765", "098", "321", "654", "987", "210", "543", "876", "109", "432",
}

var fakeBank = []string{
	"First National Bank", "Atlantic Bank", "Sunrise Bank", "Capital One Bank", "Unity Bank", "Central Bank", "Financial Trust Bank", "Pioneer Bank", "Everest Bank", "Community Bank", "Prime Bank", "First Security Bank", "Liberty Bank", "Golden State Bank", "Northwest Bank", "Southside Bank", "Harmony Bank", "Citizen Bank", "Progressive Bank", "Heartland Bank", "Crescent Bank", "Bank of the West", "City National Bank", "Heritage Bank", "Horizon Bank", "Pinnacle Bank", "Liberty National Bank", "Grand Central Bank", "Express Bank", "Master Bank", "Green Leaf Bank", "Capital Bank", "Primetime Bank", "Gateway Bank", "Union Bank", "Sovereign Bank", "Centennial Bank", "Parkside Bank", "Meridian Bank", "Federal Reserve Bank", "Interstate Bank", "Silver State Bank", "Eagle National Bank", "United Bank", "National Security Bank", "Sage Bank", "Regency Bank", "Oceanic Bank", "Lion Bank", "Cypress Bank",
}

var fakePhone = []string{
	"+1-555-123-4567", "+1-555-987-6543", "+1-555-876-5432", "+1-555-234-5678", "+1-555-789-0123",
	"+44-20-1234-5678", "+44-20-9876-5432", "+44-20-6543-2109", "+44-20-8901-2345", "+44-20-6789-0123",
	"+49-30-1234-5678", "+49-30-9876-5432", "+49-30-6543-2109", "+49-30-8901-2345", "+49-30-6789-0123",
	"+33-1-1234-5678", "+33-1-9876-5432", "+33-1-6543-2109", "+33-1-8901-2345", "+33-1-6789-0123",
	"+61-2-1234-5678", "+61-2-9876-5432", "+61-2-6543-2109", "+61-2-8901-2345", "+61-2-6789-0123",
}

var fakeEmail = []string{
	"john.smith@example.com", "emma.jones@example.com", "michael.brown@example.com", "olivia.davis@example.com", "william.clark@example.com", "ava.roberts@example.com", "james.thompson@example.com", "sophia.hall@example.com", "benjamin.white@example.com", "amelia.green@example.com", "mason.lewis@example.com", "isabella.harris@example.com", "jackson.martin@example.com", "mia.anderson@example.com", "aiden.wilson@example.com", "charlotte.taylor@example.com", "lucas.jackson@example.com", "harper.adams@example.com", "oliver.lee@example.com", "chloe.carter@example.com", "ethan.moore@example.com", "avery.hill@example.com", "logan.phillips@example.com", "lily.peterson@example.com", "sebastian.wright@example.com", "ella.cooper@example.com", "alexander.reed@example.com", "grace.morris@example.com", "daniel.murphy@example.com", "scarlett.bailey@example.com", "owen.kelly@example.com", "hannah.nelson@example.com", "wyatt.myers@example.com", "zoey.rogers@example.com", "aidan.turner@example.com", "violet.cook@example.com", "gabriel.howard@example.com", "sarah.ward@example.com", "joseph.brooks@example.com", "addison.bennett@example.com", "henry.gray@example.com", "emily.sullivan@example.com", "leo.collins@example.com", "ava.richardson@example.com", "jack.cox@example.com", "avery.bell@example.com", "mila.bailey@example.com", "ryan.wood@example.com", "layla.murphy@example.com", "nathan.harris@example.com", "evie.howard@example.com",
}

var fakeAccountNo = []string{
	"1234567890", "9876543210", "4567890123", "8901234567", "2345678901",
	"5678901234", "0123456789", "3456789012", "6789012345", "9012345678",
	"4321098765", "7654321098", "2109876543", "5432109876", "8765432109",
	"1098765432", "3210987654", "6543210987", "9876543210", "2345678901",
	"5678901234", "8901234567", "0123456789", "3456789012", "6789012345",
	"9012345678", "4321098765", "7654321098", "2109876543", "5432109876",
	"8765432109", "1098765432", "3210987654", "6543210987", "9876543210",
	"2345678901", "5678901234", "8901234567", "0123456789", "3456789012",
	"6789012345", "9012345678", "4321098765", "7654321098", "2109876543",
	"5432109876", "8765432109", "1098765432", "3210987654", "6543210987",
}

var fakeName = []string{
	"John Smith", "Emily Johnson", "Michael Williams", "Emma Jones", "Daniel Brown", "Olivia Davis", "James Miller", "Sophia Wilson", "David Taylor", "Isabella Anderson", "Joseph Martinez", "Charlotte Clark", "William Rodriguez", "Mia Lopez", "Jacob Lee", "Amelia Hernandez", "Andrew Gonzalez", "Grace Moore", "Benjamin Martin", "Sophia Young", "Joshua Lewis", "Ava Walker", "Matthew Hall", "Chloe Allen", "Christopher King", "Emily Scott", "Ethan Green", "Madison Adams", "Daniel Baker", "Abigail Nelson", "Anthony Hill", "Elizabeth Rivera", "Jackson Mitchell", "Ella Campbell", "David Carter", "Sofia Perez", "Christopher Perez", "Avery Turner", "Victoria Collins", "Logan Cooper", "Samuel Phillips", "Julia Torres", "John Stewart", "Scarlett Hughes", "Ryan Morris", "Sarah Sanchez", "Noah Price", "Grace Powell", "Jacob Bell", "Lily Ross", "Jonathan Bennett", "Alyssa Reed", "James Bailey", "Sophia Wood", "Benjamin Cook", "Ava Bailey", "Nicholas Reed", "Natalie Reed", "Alexander Parker", "Zoe Mitchell", "Mason White", "Audrey Parker", "William Wright", "Samantha Butler", "Daniel Hall", "Amelia Adams", "Joseph Turner", "Hannah Flores", "Joshua Smith", "Olivia Bennett", "Christian Lee", "Ella Hernandez", "Anthony Foster", "Ariana Campbell", "Elijah Turner", "Emily Adams", "David Cook", "Sofia Robinson", "Matthew Peterson", "Chloe Parker", "Christopher Scott", "Stella Thomas", "Michael Foster", "Lucy Martinez", "Muhammad Watson", "Evelyn Turner", "Dylan Watson", "Addison Young", "Ethan Flores", "Mia Johnson", "Alexander Cooper", "Layla Turner", "Daniel Murphy", "Jonathan Rivera", "Savannah Turner", "William Reed",
}
