package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"uhs/internal/config"

	"github.com/labstack/echo/v5"
)

// Builds Prompt for AI
func GetPrompt(jobDesc string, cvStr string) string {
	return fmt.Sprintf(`{
  "meta": {
    "role": "Senior_HR_Data_Scientist",
    "task": "CV_Analysis_and_Scoring",
    "objective": "Evaluate candidate alignment with Job Description (JD) and output strictly structured data."
  },
  "instructions": {
    "analysis_mode": "critical_and_objective",
    "output_format": "JSON_ONLY",
    "formatting_rules": [
      "Return valid JSON.",
      "No markdown formatting.",
      "No conversational filler."
    ]
  },
  "scoring_rubric": {
    "dept": {
      "type": "String",
      "description": "Primary academic or professional discipline."
    },
    "degreeLevel": {
      "type": "String",
      "description": "Highest academic degree achieved (e.g., PhD, MS, BS)."
    },
    "publicationCount": {
      "type": "Integer",
      "description": "Total number of valid research publications found."
    },
    "experience": {
      "type": "Integer (0-100)",
      "criteria": "Depth and relevance of work history relative to JD requirements."
    },
    "academicScore": {
      "type": "Integer (0-100)",
      "criteria": "Based on GPA, university prestige, and academic honors."
    },
    "researchScore": {
      "type": "Integer (0-100)",
      "criteria": "Volume of papers, citations, and relevance of research area to the job."
    },
    "teachingScore": {
      "type": "Integer (0-100)",
      "criteria": "Experience in lecturing, mentoring, or corporate training."
    },
    "industrialScore": {
      "type": "Integer (0-100)",
      "criteria": "Hands-on industry projects, corporate roles, and practical tool application."
    },
    "salaryScore": {
      "type": "Integer (0-100)",
      "criteria": "Market value alignment based on seniority. 0=Misaligned (too senior/junior), 100=Perfect fit for role level."
    },
    "adminScore": {
      "type": "Integer (0-100)",
      "criteria": "Experience in management, leadership, or administrative coordination."
    },
    "compositeRank": {
      "type": "Integer (0-100)",
      "criteria": "Weighted average heavily prioritizing skills explicitly requested in the JD."
    },
    "summary": {
      "type": "String",
      "criteria": "Concise justification of the Composite Rank (max 3 sentences)."
    }
  },
  "input_data": {
    "job_description": "%s",%s"
  }
}`, jobDesc, cvStr)

}
// Saving files on server
// Mainly files are stored in cloud like in AWS S3 Bucketss, Cloudinary etc.
// Returns the path where the file is saved
func SaveFileInServer(c *echo.Context, file *multipart.FileHeader, cfg *config.Config) (string, error) {
	src, err := file.Open()
	if err != nil {
		c.Logger().Error("Failed to open file to be saved on server " + err.Error())
		return "", err
	}
	dsn, err := os.Create(cfg.UploadsPath + "/" + file.Filename)
	if err != nil {
		c.Logger().Error("Failed to save file to server " + err.Error())
		return "", err
	}
	defer dsn.Close()

	_, err = io.Copy(dsn, src)
	if err != nil {
		c.Logger().Error("Failed to write the file to the server")
		return "", err
	}
	return cfg.UploadsPath + "/" + file.Filename, nil
}
