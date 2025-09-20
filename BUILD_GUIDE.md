# Build guide

For building Gopher ARCADE you need an iron solder and some solder. Flux is highly recommended too (specially for the SN74LVC2G04DBVR). We recommend you do it in a well ventilated room and with the proper safety precautions.

## Parts list

| No | Item name                        | Quantity | Notes |
|----|----------------------------------|------|------|
| 1  | PCB                              | 1    | 
| 2  | ATTiny85 20PU                    | 1    | 
| 3  | DIP8P socket                     | 1    | 
| 4  | SSD1306 I²C 128x64 display       | 1    | 
| 5  | 10k ohm resistors (through hole) | 2    | 
| 6  | SN74LVC2G04DBVR                  | 1    | 
| 7  | CPT-1203-78-SMT-TR (buzzer)      | 1    | 
| 8  | EG1213 switch                    | 1    | 
| 9  | CR2032 battery holder (smd)      | 1    | 
| 10 | Kailh choc v2 low profile        | 2    | footprint accepts both kailh choc v1 & v2
| 11 | Kailh choc v2 low profile keycap | 2    | 
| 12 | Display spacer (3d printed)      | 1    | optional, but helps aligning the display


## Component descriptions

### PCB

On the left of the image is the top part of the PCB and on the right is the bottom side.

## Front side
We'll use this badge to indicate the orientation during assembly: ![Work surface - front](https://img.shields.io/badge/Work_surface-front-2ea44f)

![Front side](/assets/front.jpg)


## Back side
We'll use this badge to indicate the orientation during assembly: ![Work surface - back](https://img.shields.io/badge/Work_surface-back-a42e4f)

![Back side](/assets/back.jpg)


# Soldering/assembly

The following describes soldering and assembly.
For each item, some work on the front side and some work on the back side.
The badge below indicates the orientation at the start of work, so please refer to it. 




![Work surface - front](https://img.shields.io/badge/Work_surface-front-2ea44f)

We need to prepare the board for the two SMD components, the buzzer in the middle of the board and the tiny SN74LVC2G04DBVR on the top right corner. Add a little bit of flux (if you have it) and solder on the circled areas of the image. Let's start with the SN74LVC2G04DBVR as it's the most difficult component.

![](/assets/step1.jpg)


Check the orientation, as the dot in the SN74LVC2G04DBVR should match the dot on the PCB as the following image:
![](/assets/detail.jpg)

Once the board has some solder on the pads, place the component on top of it and press it down with some tweezers. Heat one side of the SN74LVC2G04DBVR with the iron, heat the other side. Check there are no shortcircuits among the pads.


Next you can do the resitors, resistors don't have orientation.

![](/assets/step2.jpg)


Now we can do the buzzer, check the orientation, the dot should be on the top right corner. We do the same procedure as with the SN74LVC2G04DBVR, place the buzzer on top of the pre-soldered pad. Heat one side and then the other. 

![](/assets/step3.jpg)



Now on the back side we need to fit the switch on the ![back](https://img.shields.io/badge/back-a42e4f) but we solder the pins on the ![front](https://img.shields.io/badge/front-2ea44f)



![](/assets/step5.jpg)


![](/assets/step4.jpg)

![Work surface - back](https://img.shields.io/badge/Work_surface-back-a42e4f)

With the DIP8P socket is the contrary, we put it on the ![front](https://img.shields.io/badge/front-2ea44f) and solder on the ![back](https://img.shields.io/badge/back-a42e4f). Check the orientation.


![](/assets/step7.jpg)
![](/assets/step6.jpg)


Follow with the key switches. Both Kailh choc v1 or v2 are compatible. Use some clips to hold the switches if needed, or maybe some tape or some removable adhesive putty.

![](/assets/step8.jpg)
![](/assets/step9.jpg)

![Work surface - front](https://img.shields.io/badge/Work_surface-front-2ea44f)

Place the display spacer on the front as shown in the following image. Place the display too, hold it with some clips, adhesive putty or tape.

![](/assets/step10.jpg)


![Work surface - back](https://img.shields.io/badge/Work_surface-back-a42e4f)

Solder the 4 pins of the display.

![](/assets/step11.jpg)


The last step is to solder the battery holder. Add some solder to the three pads. A little bump in the middle will help the contact with the battery. Press the battery holder against the PCB and heat each one of the pads on the sides until the tin melts and keep the holder in place.

![](/assets/step12.jpg)


![Work surface - front](https://img.shields.io/badge/Work_surface-front-2ea44f)

Add the keycaps, the ATTiny85 on the socket and it's done! Congratulations on your new **Gopher ARCADE**

![Gopher ARCADE](/assets/step13.jpg)
