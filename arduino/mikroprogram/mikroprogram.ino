/*
Measuring AC Current Using ACS712

*/
#include "ACS712.h"

/*
  This example shows how to measure the power consumption
  of devices in 230V electrical system
  or any other system with alternative current
*/

// We have 30 amps version sensor connected to A0 pin of arduino
// Replace with your version if necessary
ACS712 sensor1(ACS712_05B, A1);
ACS712 sensor2(ACS712_05B, A2);
ACS712 sensor3(ACS712_05B, A3);

const int sensorIn = A0;
const int pinSensorTegangan=A1;
int mVperAmp = 100; // use 100 for 20A Module and 66 for 30A Module


double Voltage = 0;
double VRMS = 0;
double AmpsRMS = 0;

double sensorValue=0;
double sensorValue1=0;
int crosscount=0;
int climbhill=0;
double VmaxD=0;
double VeffD;
double Veff;
bool readyTosend=false;
float teganganWeb,arusWeb;
float bacaSensorArus();
float bacaSensorTegangan();
void setup(){ 
 Serial.begin(9600);
 delay(1000);
 
 Serial.println("tes");
   sensor1.calibrate();
  sensor2.calibrate();
  sensor3.calibrate();
}

void loop(){
   float U = 220;

  // To measure current we need to know the frequency of current
  // By default 50Hz is used, but you can specify own, if necessary
  float I = sensor1.getCurrentAC();
  // To calculate the power we need voltage multiplied by current
  float P = U * I;
  Serial.println(String("I = ") + I + " A");
  Serial.println(String("P = ") + P + " Watts");
 /*sensor arus*/
 Voltage = bacaSensorArus();
 VRMS = (Voltage/2.0) *0.707; 
 AmpsRMS = (VRMS * 1000)/mVperAmp;
 arusWeb=AmpsRMS;
 //Serial.print(AmpsRMS);
 //Serial.println(" Amps RMS");
 /*sensor tegangan*/

 bacaSensorTegangan();
 if(arusWeb!=0.0 && teganganWeb!=0.0){
   Serial.print(R"({"arus":")");
    Serial.print(arusWeb);
      Serial.print(R"(","tegangan":")");   
       Serial.print(teganganWeb);
              Serial.print(R"("})");
    Serial.println("#");   
  teganganWeb=0.0;
  arusWeb=0,0;
  }
 delay(500);
}
float bacaSensorTegangan(){
   //http://sentroino.blogspot.co.id/2015/12/pembacaan-tegangan-ac-menggunakan.html
sensorValue1=sensorValue;
delay(100);
sensorValue = analogRead(pinSensorTegangan);
if (sensorValue>sensorValue1 && sensorValue>511){
  climbhill=1;
  VmaxD=sensorValue;
  }
if (sensorValue<sensorValue1 && climbhill==1){
  climbhill=0;
  VmaxD=sensorValue1;
  VeffD=VmaxD/sqrt(2);
  Veff=(((VeffD-420.76)/-90.24)*-210.2)+210.2;
 // Serial.println(Veff);
 teganganWeb=Veff;
  VmaxD=0;
}
float voltageOut;
return voltageOut;
}
float bacaSensorArus()
{
  //http://henrysbench.capnfatz.com/henrys-bench/arduino-current-measurements/acs712-arduino-ac-current-tutorial/
  float result;
  
  int readValue;             //value read from the sensor
  int maxValue = 0;          // store max value here
  int minValue = 1024;          // store min value here
  uint32_t start_time = millis();
   while((millis()-start_time) < 1000) //sample for 1 Sec
   {
       readValue = analogRead(sensorIn);
       // see if you have a new maxValue
       if (readValue > maxValue) 
       {
           /*record the maximum sensor value*/
           maxValue = readValue;
       }
       if (readValue < minValue) 
       {
           /*record the maximum sensor value*/
           minValue = readValue;
       }
      
   }
   
   // Subtract min from max
   result = ((maxValue - minValue) * 5.0)/1024.0;
      
   return result;
 }
