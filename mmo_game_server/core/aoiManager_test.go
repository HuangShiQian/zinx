package core

import (
	"testing"
	"fmt"
)

func TestAOIManager_init(t *testing.T)  {
	//初始化AOIManager
	aoiMngr:=NewAOIManager(0,250,5,0,250,5)
	//打印信息
	fmt.Println(aoiMngr)
}

func TestAOIManagerSurround(t *testing.T)  {
	//初始化AOIManager
	aoiMngr:=NewAOIManager(0,250,5,0,250,5)
	//求出每个格子周边的九宫格信息
	for gid,_:=range aoiMngr.grids{
		grids:=aoiMngr.GetSurroundGridsByGid(gid)
		fmt.Println("gid : ",gid,"grids num = ",len(grids))

		//当前九宫格的ID集合
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		fmt.Println("grids IDs are ",gIDs)

	}
	fmt.Println(" |||||||||||||||||||||||||||||||  ")

	playerIDs := aoiMngr.GetSurroundPIDsByPos(175, 68)
	fmt.Println("PlayerIDs: ", playerIDs)


}