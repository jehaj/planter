import matplotlib.pyplot as plt
import numpy as np
import tensorflow as tf

import keras
from keras import layers, Sequential

import pathlib

data_dir = pathlib.Path("data").with_suffix('')
image_count = len(list(data_dir.glob('*/*.jpg')))
print(image_count)

batch_size = 4
img_height = 1280
img_width = 720

train_ds = keras.utils.image_dataset_from_directory(
  data_dir,
  validation_split=0.2,
  subset="training",
  seed=123,
  image_size=(img_height, img_width),
  batch_size=batch_size)

class_names = train_ds.class_names
print(class_names)
