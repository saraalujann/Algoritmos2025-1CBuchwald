#include <stdio.h>
#include <opencv2/opencv.hpp>
using namespace cv;

#define FLOAT_TO_INT(x) ((x)>=0?(int)((x)+0.5):(int)((x)-0.5))

int procesarImagen(uchar* p, int nRows, int nCols, int channels);
double valorRGBlineal (double RGBcomprimido);
double valorYcomprimido (double valorYlineal);


int main (int argc, char** argv) {
	int resultado;
	uchar* p;
	int channels;
	int nRows;
	int nCols;

	if (argc != 2) {
		printf ("Uso:  testopencv <imagen>\n");
		return (-1);
	}
	Mat image;
	image = imread (argv[1], IMREAD_COLOR);
	if (!image.data) {
		printf ("Sin datos de imagen... \n");
		return (-1);
	}
	namedWindow ("Original", WINDOW_AUTOSIZE);
	imshow ("Original", image);

	channels = image.channels();
	nRows = image.rows;
	nCols = image.cols;
	p = image.data;

	printf ("CTRL+C para finalizar\n\n");
	printf ("Filas: %d\n", nRows);
	printf ("Columnas: %d\n", nCols);
	printf ("Canales: %d\n", channels);
	CV_Assert(image.depth() == CV_8U);

	// Procesamiento
	resultado = procesarImagen (p, nRows, nCols, channels);

	namedWindow ("Grayscale", WINDOW_AUTOSIZE);
	imshow ("Grayscale", image);

	waitKey(0);
	return (0);
}

int procesarImagen(uchar* p, int nRows, int nCols, int channels) {

	int i,j;
	double valorR, valorG, valorB; // valores RGB en rango 0..1
	double Rlineal, Glineal, Blineal; // El array esta ordenado como BGR
	double Ylineal, Yrgb;

	for (i=0; i < nRows; i++) {
		for (j=0; j < nCols*channels; j+=3) {
			// B
			valorB = (((double)(*(p+j+i*nCols*channels)))/255.0);
			Blineal = valorRGBlineal (valorB);
			// G
			valorG = (((double)(*(p+j+i*nCols*channels+1)))/255.0);
			Glineal = valorRGBlineal (valorG);
			// R
			valorR = (((double)(*(p+j+i*nCols*channels+2)))/255.0);
			Rlineal = valorRGBlineal (valorR);
			// Y lineal
			Ylineal = 0.2126*Rlineal + 0.7152*Glineal + 0.0722*Blineal;
			// Y comprimido
			Yrgb = 255 * valorYcomprimido (Ylineal);
			// RGB grayscale
			*(p+j+i*nCols*channels) = FLOAT_TO_INT((Yrgb));
			*(p+j+i*nCols*channels+1) = FLOAT_TO_INT((Yrgb));
			*(p+j+i*nCols*channels+2) = FLOAT_TO_INT((Yrgb));
		}
	}
	return (0);
}

double valorRGBlineal (double RGBcomprimido) {
	double resultado;
	double a,b;
	if (isless(RGBcomprimido,0.04045)) {
		resultado = RGBcomprimido / 12.92;
	} else {
		a = (RGBcomprimido+0.055);
		b = (a/1.055);
		resultado = pow(b,(2.4));
	}
	return (resultado);
}

double valorYcomprimido (double valorYlineal) {
	double resultado;
	double a, b;
	if (isless(valorYlineal,0.0031308)) {
		resultado = valorYlineal * 12.92;
	} else {
		a = pow(valorYlineal,(1/2.4));
		b = 1.055 * a;
		resultado =  b - 0.055;
	}
	return (resultado);
}

