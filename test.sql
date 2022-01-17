--name:test1
Select * from test

--name: test2 
--var:id
--var:name
--var_type:int
--var:age=14
--var:text=Hola como estas 1 = 1
Select * 
from test
where 
        id = "${{id}}" 
    and name = "${{name}}" 
    and age = ${{age}} 
    and text = "${{text}}"