package controller
import (
	"net/http"
	"strconv"
	"github.com/mls93/crosszeros/model"
)


func GetTableHandler(w http.ResponseWriter, r *http.Request){

	client := GetCookie(r,CLIENT_COOKIE)
	enemy := GetCookie(r,OPPOSITE_COOKIE)

	field := model.GetPlayerField(client)
	isTurn := model.GetPlayerTurn(client)
	result := drawTable(enemy,isTurn,field)
	w.Write([]byte(result))
}


func drawTable(enemy string, isTurn bool,fld_wrap *model.Field ) string {
	field := fld_wrap.Field
	turn_string :="Your turn"
	if (!isTurn){
		turn_string = "Your opposite`s turn"
	}
	result:= "<div>Your enemy:"+enemy+"</div><div id='which_turn'>"+turn_string+"</div><div id='won_cond'></div><div id='quit_button'><a href=\"#\">Quit</a></div><table><tbody>";

	for i:=0;i<fld_wrap.Height;i++{
		result+="<tr>"
		for j:=0;j<fld_wrap.Width;j++{
			k := strconv.Itoa(i)
			l := strconv.Itoa(j)

			if (field[i][j]<1){
				result += "<td class='table_rows active-rows' id='c"+k+"x"+l+"'></td>"
			} else {
				if (field[i][j]==1){
					result += "<td  class='table_rows' id='c"+k+"x"+l+"'>O</td>"
				} else{
					result += "<td  class='table_rows' id='c"+k+"x"+l+"'>X</td>"
				}
			}

		}
		result+="</tr>"
	}
	result+="</tbody></table>";
	return result
}
