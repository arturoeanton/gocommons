--name:test1
Select * from test


--name: test2 
--var:id
--var:name
--var:age:number=14
--var:text:string=Hola como estas 1 = 1
Select * 
from test
where 
        id = "${{id}}" 
    and name = "${{name}}" 
    and age = ${{age}} 
    and text = "${{text}}"

