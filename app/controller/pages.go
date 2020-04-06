package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//указываем путь к нужному файлу
	path := filepath.Join("public", "html", "sp.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)

        //content :=  `<html lang="ru">
	//                     <head>
	//	                       <meta charset="UTF-8">
	//	                       <title>Title</title>
	//	                   </head>
	//	                   <body>
	//	                       <h1>Приветствую тебя на стартовой странице этого сайта!</h1>
	//	                       <ul>
	//	                           <li>Первый элемент маркированного списка</li>
	//		                         <li>Второй элемент маркированного списка</li>
	//		                         <li>Третий элемент маркированного списка</li>
	//		                         <li>Четвертый элемент маркированного списка</li>
	//	                       </ul>
	//	                   </body>
	//	                   </html>`
        //создаем html-шаблон
	//tmpl, err := template.New("example").Parse(content)

	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}