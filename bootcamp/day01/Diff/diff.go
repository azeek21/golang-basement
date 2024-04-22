package diff

import (
	cakes "src/Cakes"
)

var (
	add    = "ADDED "
	remove = "REMOVED "
	change = "CHANGED "
	ing    = "ingredient "
	ct     = "cooking time "
	fur    = "for "
	cn     = "count "
	u      = "unit "
	c      = "cake "
	fm     = "from "
	to     = "to "
)

func diffCakeIngredient(src cakes.CakeIngredient, tar cakes.CakeIngredient) []string {
	res := []string{}

	if len(src.Unit) == 0 && len(tar.Unit) > 0 {
		res = append(res, add+u+tar.Unit+" "+to+ing+src.Name)
	} else if len(src.Unit) > 0 && len(tar.Unit) < 1 {
		res = append(res, remove+u+src.Unit+" "+fm+ing+src.Name)
	} else if src.Unit != tar.Unit {
		res = append(res, change+u+fur+ing+src.Name+" "+fm+src.Unit+" "+to+tar.Unit)
	}

	if len(src.Count) == 0 && len(tar.Count) > 0 {
		res = append(res, add+u+cn+tar.Count+fur+ing+src.Name)
	} else if len(src.Count) > 0 && len(tar.Count) < 1 {
		res = append(res, remove+u+cn+src.Count+fm+ing+src.Name)
	} else if src.Count != tar.Count {
		res = append(res, change+u+cn+fur+ing+fm+src.Count+to+tar.Count)
	}
	return res
}

func diffCakeIngredients(
	src []cakes.CakeIngredient,
	tar []cakes.CakeIngredient,
) []string {
	diff := make(map[string]*[2]int)
	srcMap := make(map[string]cakes.CakeIngredient)
	tarMap := make(map[string]cakes.CakeIngredient)
	res := []string{}

	for _, ingred := range src {
		srcMap[ingred.Name] = ingred
		if diff[ingred.Name] == nil {
			diff[ingred.Name] = &[2]int{0, 0}
		}
		diff[ingred.Name][0] = 1

	}

	for _, ingred := range tar {
		tarMap[ingred.Name] = ingred
		if diff[ingred.Name] == nil {
			diff[ingred.Name] = &[2]int{0, 0}
		}
		diff[ingred.Name][1] = 1
	}

	for name, src_tar := range diff {
		if src_tar[0] == 1 && src_tar[1] == 1 {
			res = append(res, diffCakeIngredient(srcMap[name], tarMap[name])...)
		} else if src_tar[0] == 1 && src_tar[1] == 0 {
			res = append(res, remove+ing+name)
		} else if src_tar[0] == 0 && src_tar[1] == 1 {
			res = append(res, add+ing+name)
		}
	}

	return res
}

func diffCake(src cakes.Cake, tar cakes.Cake) []string {
	res := []string{}
	if src.Time != tar.Time {
		res = append(res, change+ct+fur+c+src.Name+" :"+fm+src.Time+to+tar.Time)
	}
	tmp := diffCakeIngredients(src.Ingredients, tar.Ingredients)
	if len(tmp) > 0 {
		res = append(res, change+c+src.Name+" ingredients: ")
		res = append(res, tmp...)
	}
	return res
}

func DiffCakes(src *[]cakes.Cake, tar *[]cakes.Cake) []string {
	srcMap := map[string]cakes.Cake{}
	tarMap := map[string]cakes.Cake{}
	diff := make(map[string]*[2]int)
	res := []string{}

	for _, cake := range *src {
		srcMap[cake.Name] = cake
		if diff[cake.Name] == nil {
			diff[cake.Name] = &[2]int{0, 0}
		}
		diff[cake.Name][0] = 1
	}

	for _, cake := range *tar {
		tarMap[cake.Name] = cake
		if diff[cake.Name] == nil {
			diff[cake.Name] = &[2]int{0, 0}
		}
		diff[cake.Name][1] = 1
	}

	for name, src_tar := range diff {
		if src_tar[0] == 1 && src_tar[1] == 1 {
			tmp := diffCake(srcMap[name], tarMap[name])
			if len(tmp) > 0 {
				res = append(res, change+c+name+":")
				res = append(res, tmp...)
			}
		} else if src_tar[0] == 1 && src_tar[1] == 0 {
			res = append(res, remove+c+name)
		} else if src_tar[0] == 0 && src_tar[1] == 1 {
			res = append(res, add+c+name)
		}
	}

	return res
}

func DiffSnapshots(src []string, tar []string) []string {
	srcMap, tarMap, diff := map[string]string{}, map[string]string{}, map[string]*[2]int{}
	res := []string{}

	for _, line := range src {
		srcMap[line] = line
		if diff[line] == nil {
			diff[line] = &[2]int{0, 0}
		}
		diff[line][0] = 1
	}

	for _, line := range tar {
		tarMap[line] = line
		if diff[line] == nil {
			diff[line] = &[2]int{0, 0}
		}
		diff[line][1] = 1
	}

	for line, src_tar := range diff {
		if src_tar[0] == 1 && src_tar[1] == 0 {
			res = append(res, remove+line)
		} else if src_tar[0] == 0 && src_tar[1] == 1 {
			res = append(res, add+line)
		}
	}

	return res
}
