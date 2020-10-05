package controllers

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"manipulasi/models"
	"manipulasi/libraries"
)

const filePath = "./companieses.csv"

func GetAllCompanies(c *gin.Context) { 
	var OneRecord models.Companies
	var AllRecords []models.Companies

	csvFile, err := os.Open(filePath)
	libraries.CheckError(err)

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	libraries.CheckError(err)

	for _, each := range csvData[1:] {
		OneRecord.CMGUnmaskedID = each[0]
		OneRecord.CMGUnmaskedName = each[1]
		OneRecord.ClientTier = each[2]
		OneRecord.GCPStream = each[3]
		OneRecord.GCPBusiness = each[4]
		OneRecord.CMGGlobalBU = each[5]
		OneRecord.CMGSegmentName = each[6]
		OneRecord.GlobalControlPoint = each[7]
		OneRecord.GCPGeography = each[8]
		OneRecord.GlobalRelationshipManagerName = each[9]
		OneRecord.REVENUE_FY14 = each[10]
		OneRecord.REVENUE_FY15 = each[11]
		OneRecord.Deposits_EOP_FY14 = each[12]
		OneRecord.Deposits_EOP_FY15x = each[13]
		OneRecord.TotalLimits_EOP_FY14 = each[14]
		OneRecord.TotalLimits_EOP_FY15 = each[15]
		OneRecord.TotalLimits_EOP_FY15x = each[16]
		OneRecord.RWAFY15 = each[17]
		OneRecord.RWAFY14 = each[18]
		OneRecord.REVRWA_FY14 = each[19]
		OneRecord.REVRWA_FY15 = each[20]
		OneRecord.NPAT_AllocEq_FY14 = each[21]
		OneRecord.NPAT_AllocEq_FY15X = each[22]
		OneRecord.Company_Avg_Activity_FY14 = each[23]
		OneRecord.Company_Avg_Activity_FY15 = each[24]
		OneRecord.ROE_FY14 = each[25]
		OneRecord.ROE_FY15 = each[26]
		AllRecords = append(AllRecords, OneRecord) 
	}

	// Convert to JSON
	_, err = json.Marshal(AllRecords)
	libraries.CheckError(err)
	c.JSON(200, gin.H{ 
		"data": AllRecords,
	})
}

func GetOneCompanies(c *gin.Context ) {
	var OneCompanies models.Companies
	var AllRecords []models.Companies
	
	csvFile, err := os.Open(filePath)
	libraries.CheckError(err)

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	for _, each := range csvData {
		if each[0] == c.Param("id") { 
			OneCompanies.CMGUnmaskedID = each[0]
			OneCompanies.CMGUnmaskedName = each[1]
			OneCompanies.ClientTier = each[2]
			OneCompanies.GCPStream = each[3]
			OneCompanies.GCPBusiness = each[4]
			OneCompanies.CMGGlobalBU = each[5]
			OneCompanies.CMGSegmentName = each[6]
			OneCompanies.GlobalControlPoint = each[7]
			OneCompanies.GCPGeography = each[8]
			OneCompanies.GlobalRelationshipManagerName = each[9]
			OneCompanies.REVENUE_FY14 = each[10]
			OneCompanies.REVENUE_FY15 = each[11]
			OneCompanies.Deposits_EOP_FY14 = each[12]
			OneCompanies.Deposits_EOP_FY15x = each[13]
			OneCompanies.TotalLimits_EOP_FY14 = each[14]
			OneCompanies.TotalLimits_EOP_FY15 = each[15]
			OneCompanies.TotalLimits_EOP_FY15x = each[16]
			OneCompanies.RWAFY15 = each[17]
			OneCompanies.RWAFY14 = each[18]
			OneCompanies.REVRWA_FY14 = each[19]
			OneCompanies.REVRWA_FY15 = each[20]
			OneCompanies.NPAT_AllocEq_FY14 = each[21]
			OneCompanies.NPAT_AllocEq_FY15X = each[22]
			OneCompanies.Company_Avg_Activity_FY14 = each[23]
			OneCompanies.Company_Avg_Activity_FY15 = each[24]
			OneCompanies.ROE_FY14 = each[25]
			OneCompanies.ROE_FY15 = each[26]
			AllRecords = append(AllRecords, OneCompanies)
		}
	}

	// Convert to JSON
	_, err = json.Marshal(AllRecords)
	libraries.CheckError(err)
	c.JSON(200, gin.H{
		"data": AllRecords,
	})
}

func PostOneCompanies(c *gin.Context) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend) 
	libraries.CheckError(err)

	defer f.Close() 

	CMGUnmaskedID := c.PostForm("CMGUnmaskedID")
	CMGUnmaskedName := c.PostForm("CMGUnmaskedName")
	ClientTier := c.PostForm("ClientTier")
	GCPStream := c.PostForm("GCPStream")
	GCPBusiness := c.PostForm("GCPBusiness")
	CMGGlobalBU := c.PostForm("CMGGlobalBU")
	CMGSegmentName := c.PostForm("CMGSegmentName")
	GlobalControlPoint := c.PostForm("GlobalControlPoint")
	GCPGeography := c.PostForm("GCPGeography")
	GlobalRelationshipManagerName := c.PostForm("GlobalRelationshipManagerName")
	REVENUE_FY14 := c.PostForm("REVENUE_FY14")
	REVENUE_FY15 := c.PostForm("REVENUE_FY15")
	Deposits_EOP_FY14 := c.PostForm("Deposits_EOP_FY14")
	Deposits_EOP_FY15x := c.PostForm("Deposits_EOP_FY15x")
	TotalLimits_EOP_FY14 := c.PostForm("TotalLimits_EOP_FY14")
	TotalLimits_EOP_FY15 := c.PostForm("TotalLimits_EOP_FY15")
	TotalLimits_EOP_FY15x := c.PostForm("TotalLimits_EOP_FY15x")
	RWAFY15 := c.PostForm("RWAFY15")
	RWAFY14 := c.PostForm("RWAFY14")
	REVRWA_FY14 := c.PostForm("REVRWA_FY14")
	REVRWA_FY15 := c.PostForm("REVRWA_FY15")
	NPAT_AllocEq_FY14 := c.PostForm("NPAT_AllocEq_FY14")
	NPAT_AllocEq_FY15X := c.PostForm("NPAT_AllocEq_FY15X")
	Company_Avg_Activity_FY14 := c.PostForm("Company_Avg_Activity_FY14")
	Company_Avg_Activity_FY15 := c.PostForm("Company_Avg_Activity_FY15")
	ROE_FY14 := c.PostForm("ROE_FY14")
	ROE_FY15 := c.PostForm("ROE_FY15")

	var data [][]string
	data = append(data, []string{ 
		CMGUnmaskedID,
		CMGUnmaskedName,
		ClientTier,
		GCPStream,
		GCPBusiness,
		CMGGlobalBU,
		CMGSegmentName,
		GlobalControlPoint,
		GCPGeography,
		GlobalRelationshipManagerName,
		REVENUE_FY14,
		REVENUE_FY15,
		Deposits_EOP_FY14,
		Deposits_EOP_FY15x,
		TotalLimits_EOP_FY14,
		TotalLimits_EOP_FY15,
		TotalLimits_EOP_FY15x,
		RWAFY15,
		RWAFY14,
		REVRWA_FY14,
		REVRWA_FY15,
		NPAT_AllocEq_FY14,
		NPAT_AllocEq_FY15X,
		Company_Avg_Activity_FY14,
		Company_Avg_Activity_FY15,
		ROE_FY14,
		ROE_FY15,
	})

	w := csv.NewWriter(f) 
	w.WriteAll(data) 

	libraries.CheckError(err)
	c.JSON(200, gin.H{
		"data": map[string]string{
			"CMGUnmaskedID":                 CMGUnmaskedID,
			"CMGUnmaskedName":               CMGUnmaskedName,
			"ClientTier":                    ClientTier,
			"GCPStream":                     GCPStream,
			"GCPBusiness":                   GCPBusiness,
			"CMGGlobalBU":                   CMGGlobalBU,
			"CMGSegmentName":                CMGSegmentName,
			"GlobalControlPoint":            GlobalControlPoint,
			"GCPGeography":                  GCPGeography,
			"GlobalRelationshipManagerName": GlobalRelationshipManagerName,
			"REVENUE_FY14":                  REVENUE_FY14,
			"REVENUE_FY15":                  REVENUE_FY15,
			"Deposits_EOP_FY14":             Deposits_EOP_FY14,
			"Deposits_EOP_FY15x":            Deposits_EOP_FY15x,
			"TotalLimits_EOP_FY14":          TotalLimits_EOP_FY14,
			"TotalLimits_EOP_FY15":          TotalLimits_EOP_FY15,
			"TotalLimits_EOP_FY15x":         TotalLimits_EOP_FY15x,
			"RWAFY15":                       RWAFY15,
			"RWAFY14":                       RWAFY14,
			"REVRWA_FY14":                   REVRWA_FY14,
			"REVRWA_FY15":                   REVRWA_FY15,
			"NPAT_AllocEq_FY14":             NPAT_AllocEq_FY14,
			"NPAT_AllocEq_FY15X":            NPAT_AllocEq_FY15X,
			"Company_Avg_Activity_FY14":     Company_Avg_Activity_FY14,
			"Company_Avg_Activity_FY15":     Company_Avg_Activity_FY15,
			"ROE_FY14":                      ROE_FY14,
			"ROE_FY15":                      ROE_FY15,
		},
	})
}

func UpdateOneCompanies(c *gin.Context) {

	//open file
	f, err := os.Open(filePath)
	libraries.CheckError(err)

	rows, err := csv.NewReader(f).ReadAll() 

	f.Close() 
	libraries.CheckError(err)

	CMGUnmaskedID := c.PostForm("CMGUnmaskedID")
	CMGUnmaskedName := c.PostForm("CMGUnmaskedName")
	ClientTier := c.PostForm("ClientTier")
	GCPStream := c.PostForm("GCPStream")
	GCPBusiness := c.PostForm("GCPBusiness")
	CMGGlobalBU := c.PostForm("CMGGlobalBU")
	CMGSegmentName := c.PostForm("CMGSegmentName")
	GlobalControlPoint := c.PostForm("GlobalControlPoint")
	GCPGeography := c.PostForm("GCPGeography")
	GlobalRelationshipManagerName := c.PostForm("GlobalRelationshipManagerName")
	REVENUE_FY14 := c.PostForm("REVENUE_FY14")
	REVENUE_FY15 := c.PostForm("REVENUE_FY15")
	Deposits_EOP_FY14 := c.PostForm("Deposits_EOP_FY14")
	Deposits_EOP_FY15x := c.PostForm("Deposits_EOP_FY15x")
	TotalLimits_EOP_FY14 := c.PostForm("TotalLimits_EOP_FY14")
	TotalLimits_EOP_FY15 := c.PostForm("TotalLimits_EOP_FY15")
	TotalLimits_EOP_FY15x := c.PostForm("TotalLimits_EOP_FY15x")
	RWAFY15 := c.PostForm("RWAFY15")
	RWAFY14 := c.PostForm("RWAFY14")
	REVRWA_FY14 := c.PostForm("REVRWA_FY14")
	REVRWA_FY15 := c.PostForm("REVRWA_FY15")
	NPAT_AllocEq_FY14 := c.PostForm("NPAT_AllocEq_FY14")
	NPAT_AllocEq_FY15X := c.PostForm("NPAT_AllocEq_FY15X")
	Company_Avg_Activity_FY14 := c.PostForm("Company_Avg_Activity_FY14")
	Company_Avg_Activity_FY15 := c.PostForm("Company_Avg_Activity_FY15")
	ROE_FY14 := c.PostForm("ROE_FY14")
	ROE_FY15 := c.PostForm("ROE_FY15")

	// loop data
	var id bool = false
	for i := 1; i < len(rows); i++ { 
		if rows[i][0] == c.PostForm("CMGUnmaskedID") { 
			rows[i][1] = CMGUnmaskedName
			rows[i][2] = ClientTier
			rows[i][3] = GCPStream
			rows[i][4] = GCPBusiness
			rows[i][5] = CMGGlobalBU
			rows[i][6] = CMGSegmentName
			rows[i][7] = GlobalControlPoint
			rows[i][8] = GCPGeography
			rows[i][9] = GlobalRelationshipManagerName
			rows[i][10] = REVENUE_FY14
			rows[i][11] = REVENUE_FY15
			rows[i][12] = Deposits_EOP_FY14
			rows[i][13] = Deposits_EOP_FY15x
			rows[i][14] = TotalLimits_EOP_FY14
			rows[i][15] = TotalLimits_EOP_FY15
			rows[i][16] = TotalLimits_EOP_FY15x
			rows[i][17] = RWAFY15
			rows[i][18] = RWAFY14
			rows[i][19] = REVRWA_FY14
			rows[i][20] = REVRWA_FY15
			rows[i][21] = NPAT_AllocEq_FY14
			rows[i][22] = NPAT_AllocEq_FY15X
			rows[i][23] = Company_Avg_Activity_FY14
			rows[i][24] = Company_Avg_Activity_FY15
			rows[i][25] = ROE_FY14
			rows[i][26] = ROE_FY15
			id = true
		}
	}

	// create file again
	f, err = os.Create(filePath) 
	libraries.CheckError(err)

	err = csv.NewWriter(f).WriteAll(rows) 
	libraries.CheckError(err)

	if id {
		c.JSON(200, gin.H{
			"msg": "sukses update data dengan id : " + CMGUnmaskedID,
		})
	} else {
		c.JSON(500, gin.H{
			"msg": "data tidak ditemukan",
		})
	}
}

// library udpate

func sum(row []string) string { 
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s) 
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum) 
}