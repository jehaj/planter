from os.path import exists

import matplotlib.pyplot as plt
import numpy as np
import tensorflow as tf

import keras
from keras import layers, Sequential

import pathlib

batch_size = 4
img_height = 144
img_width = 256
seed = 123

model_path = "fitted_model_smaller"


def train(epochs=10):
    data_dir = pathlib.Path("data").with_suffix('')
    image_count = len(list(data_dir.glob('*/*.jpg')))
    print(image_count)
    train_ds = keras.utils.image_dataset_from_directory(
        data_dir,
        validation_split=0.2,
        subset="training",
        seed=seed,
        image_size=(img_height, img_width),
        batch_size=batch_size)
    val_ds = keras.utils.image_dataset_from_directory(
        data_dir,
        validation_split=0.2,
        subset="validation",
        seed=seed,
        image_size=(img_height, img_width),
        batch_size=batch_size)

    class_names = train_ds.class_names

    # AUTOTUNE = tf.data.AUTOTUNE
    #
    # train_ds = train_ds.cache().shuffle(1000).prefetch(buffer_size=AUTOTUNE)
    # val_ds = val_ds.cache().prefetch(buffer_size=AUTOTUNE)

    normalization_layer = layers.Rescaling(1. / 255)

    normalized_ds = train_ds.map(lambda x, y: (normalization_layer(x), y))
    image_batch, labels_batch = next(iter(normalized_ds))
    first_image = image_batch[0]

    num_classes = len(class_names)

    model = Sequential([
        layers.Rescaling(1. / 255, input_shape=(img_height, img_width, 3)),
        layers.Conv2D(16, 3, padding='same', activation='relu'),
        layers.MaxPooling2D(),
        layers.Conv2D(32, 3, padding='same', activation='relu'),
        layers.MaxPooling2D(),
        layers.Conv2D(64, 3, padding='same', activation='relu'),
        layers.MaxPooling2D(),
        layers.Flatten(),
        layers.Dense(128, activation='relu'),
        layers.Dense(num_classes)
    ])

    model.compile(optimizer='adam',
                  loss=keras.losses.SparseCategoricalCrossentropy(from_logits=True),
                  metrics=['accuracy'])

    model.summary()

    history = model.fit(
        train_ds,
        validation_data=val_ds,
        epochs=epochs
    )

    model.export(model_path)

    acc = history.history['accuracy']
    val_acc = history.history['val_accuracy']

    loss = history.history['loss']
    val_loss = history.history['val_loss']

    epochs_range = range(epochs)

    plt.figure(figsize=(8, 8))
    plt.subplot(1, 2, 1)
    plt.plot(epochs_range, acc, label='Training Accuracy')
    plt.plot(epochs_range, val_acc, label='Validation Accuracy')
    plt.legend(loc='lower right')
    plt.title('Training and Validation Accuracy')

    plt.subplot(1, 2, 2)
    plt.plot(epochs_range, loss, label='Training Loss')
    plt.plot(epochs_range, val_loss, label='Validation Loss')
    plt.legend(loc='upper right')
    plt.title('Training and Validation Loss')
    plt.show()


def test():
    model = keras.layers.TFSMLayer(model_path, call_endpoint='serve')
    img = keras.utils.load_img(
        "test_data/rulle.jpg", target_size=(img_height, img_width)
    )
    img_array = keras.utils.img_to_array(img)
    img_array = tf.expand_dims(img_array, 0)  # Create a batch

    predictions = model.call(img_array)
    score = tf.nn.softmax(predictions[0])

    print(
        "This image most likely belongs to {} with a {:.2f} percent confidence."
        .format(["lampe", "rulle"][np.argmax(score)], 100 * np.max(score))
    )


def quantize():
    converter = tf.lite.TFLiteConverter.from_saved_model(model_path)
    converter.optimizations = [tf.lite.Optimize.DEFAULT]
    tflite_model = converter.convert()
    with open("model_quantized.tflite", "wb") as f:
        f.write(tflite_model)


def test_quantized_model():
    interpreter = tf.lite.Interpreter(model_path="model_quantized.tflite")
    interpreter.allocate_tensors()

    input_details = interpreter.get_input_details()
    img = keras.utils.load_img(
        "test_data/rulle.jpg", target_size=(img_height, img_width), keep_aspect_ratio=True,
    )
    input_data = tf.image.convert_image_dtype(img, dtype=tf.float32)
    input_data = tf.expand_dims(input_data, axis=0)
    interpreter.set_tensor(input_details[0]['index'], input_data)

    interpreter.invoke()

    output_details = interpreter.get_output_details()
    output_data = interpreter.get_tensor(output_details[0]['index'])

    predicted_class = np.argmax(output_data)
    confidence = output_data[0][predicted_class] # should probably use softmax to get between [0, 1]

    print(
        "This image most likely belongs to {} with a {:.2f} confidence."
        .format(["lampe", "rulle"][predicted_class], confidence)
    )


if __name__ == '__main__':
    if input("Train model? (y/N): ").lower() == "y":
        train(5)
        if input("Quantize model? (y/N): ").lower() == "y":
            quantize()
    if exists(model_path):
        test()
    if exists("model_quantized.tflite"):
        test_quantized_model()
