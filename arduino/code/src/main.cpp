#include <Arduino.h>

#include "ACS712.h"
#include <ArduinoJson.h>



// init pin acs
ACS712 sensor1(ACS712_05B, A1);
ACS712 sensor2(ACS712_05B, A2);
ACS712 sensor3(ACS712_05B, A3);
/****** var for reading current sensor*****/
 double sensorValue=0;
double sensorValue1=0;
int crosscount=0;
int climbhill=0;
double VmaxD=0;
double VeffD;
double Veff;
float bacaSensorTegangan(){
   //http://sentroino.blogspot.co.id/2015/12/pembacaan-tegangan-ac-menggunakan.html
sensorValue1=sensorValue;
delay(100);
sensorValue = analogRead(A0);
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
  VmaxD=0;
  return Veff;
}
return -1;
}
// struct data JSON
struct dataJSONs {
float tegangan;
float arus1;
float arus2;
float arus3;
  
} dataJSON;  
 StaticJsonBuffer<200> jsonBuffer;
   JsonObject& root =jsonBuffer.createObject();
void sendJSON(){
root["tegangan"] = String(dataJSON.tegangan);
root["arus1"] = String(dataJSON.arus1);
root["arus2"] = String(dataJSON.arus2);
root["arus3"] = String(dataJSON.arus3);
}
void setup() {
  Serial.begin(9600);

  // calibrate() method calibrates zero point of sensor
sensor1.calibrate();
  sensor2.calibrate();
  sensor3.calibrate();
//   Serial.println("Done!");
}

void loop() {
// baca tegangan sensor , jika sukses nilai != -1
  float U =bacaSensorTegangan();
  if (U!=-1){
dataJSON.tegangan=U;
// dapatkan 3 nilai sensor
dataJSON.arus1=sensor1.getCurrentAC();
dataJSON.arus2=sensor2.getCurrentAC();
dataJSON.arus3=sensor3.getCurrentAC();
  sendJSON();
  delay(3000);
  }
}