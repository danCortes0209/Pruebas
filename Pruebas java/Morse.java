import javax.swing.JOptionPane;

public class Morse {
	public static int incremento;
	public static String morse;

	public static void main(String[] args) {
		cifrar();
	}

	public static void cifrar(){
		String message = JOptionPane.showInputDialog("Cifrador en codigo morse (solo se aceptan numeros y letras)\nDeme un mensaje:",""); 
    	String variable = "";
        for (int x=0;x<message.length();x++){
            incremento = (int) message.codePointAt(x);
            caracteres();
            variable = "" + variable + morse;
        }
        message = variable;
		JOptionPane.showMessageDialog(null,"Mensaje cifrado: "+message);
	}

	public static void caracteres(){
		if(incremento==32){//Space
			morse="/";
		}else if(incremento==48){//0
			morse="----- ";
		}else if(incremento==49){//1
			morse=".---- ";
		}else if (incremento==50) {//2
			morse="..--- ";
		}else if (incremento==51) {//3
			morse="...-- ";
		}else if (incremento==52) {//4
			morse="....- ";
		}else if (incremento==53) {//5
			morse="..... ";
		}else if (incremento==54) {//6
			morse="-.... ";
		}else if (incremento==55) {//7
			morse="--... ";
		}else if (incremento==56) {//8
			morse="---.. ";
		}else if (incremento==57) {//9
			morse="----. ";
		}else if (incremento==65||incremento==97) {//a
			morse=".- ";
		}else if (incremento==66||incremento==98) {//b
			morse="-... ";
		}else if (incremento==67||incremento==99) {//c
			morse="-.-. ";
		}else if (incremento==68||incremento==100) {//d
			morse="-.. ";
		}else if (incremento==69||incremento==101) {//e
			morse=". ";
		}else if (incremento==70||incremento==102) {//f
			morse="..-. ";
		}else if (incremento==71||incremento==103) {//g
			morse="--. ";
		}else if (incremento==72||incremento==104) {//h
			morse=".... ";
		}else if (incremento==73||incremento==105) {//i
			morse=".. ";
		}else if (incremento==74||incremento==106) {//j
			morse=".--- ";
		}else if (incremento==75||incremento==107) {//k
			morse="-.- ";
		}else if (incremento==76||incremento==108) {//l
			morse=".-.. ";
		}else if (incremento==77||incremento==109) {//m
			morse="-- ";
		}else if (incremento==78||incremento==110) {//n
			morse="-. ";
		}else if (incremento==79||incremento==111) {//o
			morse="--- ";
		}else if (incremento==80||incremento==112) {//p
			morse=".--. ";
		}else if (incremento==81||incremento==113) {//q
			morse="--.- ";
		}else if (incremento==82||incremento==114) {//r
			morse=".-. ";
		}else if (incremento==83||incremento==115) {//s
			morse="... ";
		}else if (incremento==84||incremento==116) {//t
			morse="- ";
		}else if (incremento==85||incremento==117) {//u
			morse="..- ";
		}else if (incremento==86||incremento==118) {//v
			morse="...- ";
		}else if (incremento==87||incremento==119) {//w
			morse=".-- ";
		}else if (incremento==88||incremento==120) {//x
			morse="-..- ";
		}else if (incremento==89||incremento==121) {//y
			morse="-.-- ";
		}else if (incremento==90||incremento==122) {//z
			morse="--.. ";
		}
	}
}