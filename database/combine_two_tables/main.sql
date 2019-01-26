select
    p.FirstName as FirstName,
    p.LastName as LastName,
    Address.City as City,
    Address.State as State
from Person as p LEFT JOIN Address ON Address.PersonId = p.PersonId
