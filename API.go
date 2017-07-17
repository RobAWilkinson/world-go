package main
import	(
  "database/sql"
  "fmt"
  "strconv"
)

func GetCity(db *sql.DB, id string) City {
  var data City
  query := fmt.Sprint("SELECT * FROM CITY WHERE Id='", id, "'")
  fmt.Println(query)
  rows, err := db.Query(query)
  if( err != nil) {
    panic(err.Error())
  }
  defer rows.Close()
  for rows.Next() {
    fmt.Println(rows.Columns())
    err := rows.Scan(&data.ID, &data.Name, &data.CountryCode, &data.District, &data.Population)
    if err != nil {
      panic(err.Error())
    }
    fmt.Println(data)
  }
  return data
}

func GetCities(db *sql.DB, page string) []City {
  var data []City
  pagecount, err := strconv.Atoi(page)
  if err != nil {
    panic(err.Error())
  }
  pagecount = pagecount * 50
  query := fmt.Sprint("SELECT * FROM CITY LIMIT 50 OFFSET ", pagecount)
  rows, err := db.Query(query)
  if( err != nil) {
    panic(err.Error())
  }
  defer rows.Close()
  for rows.Next() {
    var city City
    fmt.Println(rows.Columns())
    err := rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
    if err != nil {
      panic(err.Error())
    }
    data = append(data, city)
  }
  return data;
}

func GetCountry(db *sql.DB, code string) Country {
  var country Country
  query := fmt.Sprint("SELECT * FROM COUNTRY WHERE Code='", code, "'")
  rows, err := db.Query(query)
  if( err != nil) {
    panic(err.Error())
  }
  defer rows.Close()
  fmt.Println(rows.Columns())
  for rows.Next() {
    err := rows.Scan(&country.Code, &country.Name, &country.Continent, &country.Region, &country.SurfaceArea, &country.IndepYear, &country.Population, &country.LifeExpectancy, &country.GNP, &country.GNPOld, &country.LocalName, &country.GovernmentForm, &country.HeadofState, &country.Capital, &country.Code2)
    if err != nil {
      panic(err.Error())
    }
    fmt.Println(country)
  }
  return country;
}
