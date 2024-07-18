from flask import Flask, jsonify, request
app = Flask(__name__)


@app.route('/')
def home():
    return "Welcome to the API server!"


dataList = [
    {
        "name": "Apple",
        "description": "A sweet, edible fruit produced by an apple tree.",
        "nutritionalValue": "52 kcal per 100g",
        "origin": "Central Asia"
    },
    {
        "name": "Banana",
        "description": "An elongated, edible fruit produced by several kinds of large herbaceous flowering plants.",
        "nutritionalValue": "89 kcal per 100g",
        "origin": "Southeast Asia"
    },
    {
        "name": "Cherry",
        "description": "A small, round stone fruit that is typically bright or dark red.",
        "nutritionalValue": "50 kcal per 100g",
        "origin": "Europe and Asia"
    },
    {
        "name": "Date",
        "description": "The fruit of the date palm, which is a sweet and chewy fruit.",
        "nutritionalValue": "277 kcal per 100g",
        "origin": "Middle East"
    },
    {
        "name": "Elderberry",
        "description": "A small, dark-purple fruit that grows in clusters and is known for its medicinal properties.",
        "nutritionalValue": "73 kcal per 100g",
        "origin": "Europe, Africa, and parts of Asia"
    },
    {
        "name": "Fig",
        "description": "A soft fruit with a thin skin that can be eaten ripe or dried.",
        "nutritionalValue": "74 kcal per 100g",
        "origin": "Western Asia"
    },
    {
        "name": "Grape",
        "description": "A fruit, botanically a berry, of the deciduous woody vines of the flowering plant genus Vitis.",
        "nutritionalValue": "69 kcal per 100g",
        "origin": "Near East"
    },
    {
        "name": "Honeydew",
        "description": "A fruit that has a smooth, pale outer skin and sweet, green flesh inside.",
        "nutritionalValue": "36 kcal per 100g",
        "origin": "West Africa"
    },
    {
        "name": "Indian Fig",
        "description": "Also known as prickly pear, a species of cactus that produces an edible fruit.",
        "nutritionalValue": "41 kcal per 100g",
        "origin": "Mexico"
    },
    {
        "name": "Jackfruit",
        "description": "The largest fruit that grows on a tree, with a distinctive sweet and fruity aroma.",
        "nutritionalValue": "95 kcal per 100g",
        "origin": "South India"
    }
]


@app.route('/api/fruits', methods=['GET'])
def get_data():
    start_index = request.args.get('startIndex', default=0, type=int)
    end_index = request.args.get('endIndex', default=9, type=int)

    if start_index < 0 or end_index >= 10 or start_index > end_index:
        return jsonify({"error": "Invalid index range"}), 400

    result = dataList[start_index:end_index + 1]
    return jsonify(result)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=False)
