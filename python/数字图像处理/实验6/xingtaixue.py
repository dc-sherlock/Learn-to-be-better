import numpy as np 
import cv2
img = cv2.imread('E:/python_opencv_test/j.jpg',0)
img1=cv2.imread('E:/python_opencv_test/bi_j.png',0)
img2=cv2.imread('E:/python_opencv_test/kai_j.jpg',0)
kernel = np.ones((5,5),np.uint8)
erosion = cv2.erode(img,kernel,iterations = 1)
dilation = cv2.dilate(img,kernel,iterations = 1)
opening = cv2.morphologyEx(img1, cv2.MORPH_OPEN, kernel)
closing = cv2.morphologyEx(img2, cv2.MORPH_CLOSE, kernel)
gradient = cv2.morphologyEx(img, cv2.MORPH_GRADIENT, kernel)
cv2.imshow('original',img)
cv2.imshow('erosion',erosion)
cv2.imshow('dilation',dilation)
cv2.imshow('original1',img1)
cv2.imshow('opening',opening)
cv2.imshow('original2',img2)
cv2.imshow('closing',closing)
cv2.imshow('gradient',gradient)
cv2.waitKey(0)
cv2.destroyAllWindows()

