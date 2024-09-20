import http.server
import json
import socketserver
import os
import tensorflow as tf
import numpy as np
import keras
from keras.models import load_model

PORT = 8000
UPLOAD_DIR = "upload"

# Ensure the upload directory exists
os.makedirs(UPLOAD_DIR, exist_ok=True)
model = load_model("Coral_Recog_Model.keras")


def classify_images(image_path):
    coral_types = [
        "Black Band Disease",
        "Bleached ",
        "Healthy",
        "Porites Trematodiasis",
    ]
    input_image = tf.keras.utils.load_img(image_path, target_size=(256, 256))
    input_image_array = tf.keras.utils.img_to_array(input_image)
    input_image_exp_dim = tf.expand_dims(input_image_array, 0)

    predictions = model.predict(input_image_exp_dim)
    result = tf.nn.softmax(predictions[0])
    outcome = coral_types[np.argmax(result)]

    print(outcome)

    return outcome


class MyHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/ai/classify":
            self.send_response(200)
            self.send_header("Content-type", "text/plain")
            self.end_headers()

            outcome = classify_images("./upload/uploaded_image").encode()

            self.wfile.write(outcome)
        else:
            # Respond with a 404 if any other path is requested
            self.send_response(404)
            self.send_header("Content-type", "text/plain")
            self.end_headers()
            self.wfile.write(b"404 Not Found")


with socketserver.TCPServer(("", PORT), MyHandler) as httpd:
    print(f"Serving at port {PORT}")
    httpd.serve_forever()
