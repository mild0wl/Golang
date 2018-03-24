package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"go/build"
	"math/rand"
	"os"
)

func main() {

	var tree base.Classifier

	rand.Seed(44111342)

	dataset := fmt.Sprintf("%s/src/github.com/sjwhitworth/golearn/examples/datasets/iris_headers.csv", build.Default.GOPATH)

	// Load in the iris dataset
	iris, err := base.ParseCSVToInstances(dataset, true)
	if err != nil {
		panic(err)
	}

	// Discretise the iris dataset with Chi-Merge
	filt := filters.NewChiMergeFilter(iris, 0.999)
	for _, a := range base.NonClassFloatAttributes(iris) {
		filt.AddAttribute(a)
	}
	filt.Train()
	irisf := base.NewLazilyFilteredInstances(iris, filt)

	// Create a 60-40 training-test split
	trainData, testData := base.InstancesTrainTestSplit(irisf, 0.60)

	fmt.Println(testData)

	//
	// Finally, Random Forests
	//
	tree = ensemble.NewRandomForest(70, 3)
	err = tree.Fit(trainData)
	if err != nil {
		panic(err)
	}
	predictions, err := tree.Predict(testData)
	if err != nil {
		panic(err)
	}
	fmt.Println("RandomForest Performance")
	cf, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(cf))

	modelFile := fmt.Sprintf("%sgolearn_mdl", os.TempDir())
	fmt.Printf("Saving to %s\n", modelFile)
	err = tree.Save(modelFile)
	if err != nil {
		panic(err)
	}

	var newRf base.Classifier
	newRf = ensemble.NewRandomForest(70, 3)
	newRf.Load(modelFile)
	fmt.Println(newRf)

	inst := base.NewDenseInstances()

	attrs := []string{"Sepal length", "Sepal width", "Petal length", "Petal width", "Species"}
	for _, a := range attrs {
		n := new(base.CategoricalAttribute)
		n.SetName(a)
		inst.AddAttribute(n)
	}

	attrSpecs := make([]base.AttributeSpec, len(attrs))
	attrSpecsUnordered := base.ResolveAllAttributes(inst)
	for _, a := range attrSpecsUnordered {
		name := a.GetAttribute().GetName()
		for i, b := range attrs {
			if name == b {
				attrSpecs[i] = a
			}
		}
	}

	inst.Extend(1)

	fmt.Println(inst.AllAttributes())

	inst.AddClassAttribute(inst.AllAttributes()[4])

	inst.Set(attrSpecs[0], 0, attrSpecs[0].GetAttribute().GetSysValFromString("5.100000"))
	inst.Set(attrSpecs[1], 0, attrSpecs[1].GetAttribute().GetSysValFromString("3.500000"))
	inst.Set(attrSpecs[2], 0, attrSpecs[2].GetAttribute().GetSysValFromString("1.400000"))
	inst.Set(attrSpecs[3], 0, attrSpecs[3].GetAttribute().GetSysValFromString("0.200000"))

	// This can't be unset, and whatever value it's set to becomes the prediction
	inst.Set(attrSpecs[4], 0, attrSpecs[4].GetAttribute().GetSysValFromString("test"))

	fmt.Println(inst)

	predictionsNew, err := newRf.Predict(inst)
	if err != nil {
		panic(err)
	}

	// Expected output: value in [Iris-setosa Iris-versicolor Iris-virginica]
	// Actual output: test
	fmt.Printf("\n\nPrediction: %s\n", predictionsNew.RowString(0))

}
